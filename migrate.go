// Copyright 2016 Muhammad Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE file.

package j2p

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
	"github.com/shuLhan/gonduit"
)

//
// MigrateProjects will migrate all projects from JIRA to Phabricator.
// Here is what this function will do,
// (1) It will get all project from JIRA
// (2) For each project in JIRA
// (2.1) Create the project in Phabricator
// (2.2) If its fails -- the project already exist, get the project PHID from
// Phabricator
// (2.3) Save the first matched project as mapping for Phabricator project,
// since there is no way we can check other than name.
//
func (cmd *Cmd) MigrateProjects() (e error) {
	var jiraProjects *[]jira.Project

	// 1
	jiraProjects, e = cmd.JiraGetProjects()
	if e != nil {
		return e
	}

	cmd.gonProjects = make([]gonduit.Project, len(*jiraProjects))

	// 2
	for x, project := range *jiraProjects {
		if DEBUG >= 1 {
			fmt.Printf("[j2p] MigrateProjects %d >> %s\n", x,
				project.Name)
		}

		cmd.gonProjects[x].SetName(project.Name)

		// 2.1
		e = cmd.gonProjects[x].Create(cmd.GonduitCl)

		if e == nil {
			if DEBUG >= 1 {
				fmt.Printf("[j2p] MigrateProjects << %v\n",
					cmd.gonProjects[x])
			}
			continue
		}

		// 2.2
		matchedProjects, _ := cmd.GonduitCl.ProjectSearchByName(
			project.Name)

		// 2.3
		if len(matchedProjects) == 0 {
			fmt.Printf("[j2p] MigrateProjects << "+
				"No project name matched with '%s'",
				project.Name)
			continue
		}

		cmd.gonProjects[x] = matchedProjects[0]

		if DEBUG >= 1 {
			fmt.Printf("[j2p] MigrateProjects << %v\n",
				cmd.gonProjects[x])
		}
	}

	return nil
}

func debugJiraIssue(x int, issue jira.Issue) {
	fmt.Printf("\n>>\n[j2p] migrateJiraIssues %d >>\n", x)
	fmt.Printf("    Summary >> %s\n", issue.Fields.Summary)
	fmt.Printf("    Priority >> %s\n", issue.Fields.Priority.Name)
	fmt.Printf("    Status >> %s\n", issue.Fields.Status.Name)
	fmt.Printf("    Description >> %s\n", issue.Fields.Description)

	if issue.Fields.Comments == nil {
		return
	}
	for _, c := range issue.Fields.Comments.Comments {
		if c == nil {
			continue
		}
		fmt.Printf("    Comment >> %s\n", c.Body)
	}
}

//
// migrateJiraIssues given a list of JIRA `issues` and project `pr` in
// Phabricator,
// (1) For each issue in list,
// (1.1) Convert JIRA issue to Maniphest Task
// (1.2) Search similar task in Maniphest tagged with project `pr`
// (1.3) If similar task found, skip it
// (1.4) Otherwise create new task in Maniphest
//
func (cmd *Cmd) migrateJiraIssues(
	jiraIssues []jira.Issue,
	pr gonduit.Project,
) {
	var e error
	var exist bool

	task := gonduit.Task{}

	for x, issue := range jiraIssues {
		exist = cmd.GonduitCl.TaskIsExist("all", issue.Fields.Summary,
			[]string{
				pr.GetName(),
			})

		if exist {
			fmt.Printf("[j2p] Task exist >> %s\n",
				issue.Fields.Summary)
			continue
		}

		if DEBUG >= 1 {
			debugJiraIssue(x, issue)
		}

		task.SetName(issue.Fields.Summary)
		task.SetDescription(issue.Fields.Description)
		task.SetPriority(cmd.Config.PrioritiesMap[issue.Fields.Priority.Name])
		task.SetStatus(cmd.Config.StatusesMap[issue.Fields.Status.Name])
		task.AddProject(pr.GetPhid())

		if issue.Fields.Comments != nil {
			for _, comment := range issue.Fields.Comments.Comments {
				task.AddComment(comment.Body)
			}
		}

		if DEBUG >= 2 {
			fmt.Printf("==\n[j2p] migrateJiraIssues TASK >> %+v\n",
				task)
		}

		e = task.Create(cmd.GonduitCl)

		if e != nil {
			fmt.Printf("[j2p] migrateJiraIssues ERROR: %v\n", e)
		}
	}
}

//
// MigrateTasks will migrate all task from JIRA to Phabricator based on list of
// project registered in `cmd.gonProjects`.
//
func (cmd *Cmd) MigrateTasks() (e error) {
	var q string
	var jiraIssues []jira.Issue
	var resp *jira.Response

	e = cmd.MigrateProjects()
	if e != nil {
		return
	}

	var searchOpt = &jira.SearchOptions{
		StartAt:    cmd.Config.Jira.SearchOpt.StartAt,
		MaxResults: cmd.Config.Jira.SearchOpt.MaxResults,
	}

	for _, project := range cmd.gonProjects {
		// Skip project that does not have Phid.
		if project.GetPhid() == "" {
			continue
		}
		if project.GetName() != "Proofn Dev/Ops" {
			continue
		}

		q = fmt.Sprintf(JiraQIssueAll, project.GetName())

		if DEBUG >= 1 {
			fmt.Printf("[j2p] MigrateTasks >> %s\n", q)
		}

		jiraIssues, resp, e = cmd.JiraCl.Issue.Search(q, searchOpt)

		if e != nil {
			fmt.Printf("[j2p] MigrateTasks ERROR >> %v", resp)
			return
		}

		cmd.migrateJiraIssues(jiraIssues, project)
	}

	return nil
}
