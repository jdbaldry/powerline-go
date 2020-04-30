package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
	"strings"
)

func segmentGitLite(p *powerline) []pwl.Segment {
	git := gitCommand(p.cwd)

	if len(p.ignoreRepos) > 0 {
		out, err := runGitCommand(git, "rev-parse", "--show-toplevel")
		if err != nil {
			return []pwl.Segment{}
		}
		out = strings.TrimSpace(out)
		if p.ignoreRepos[out] {
			return []pwl.Segment{}
		}
	}

	out, err := runGitCommand(git, "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return []pwl.Segment{}
	}

	status := strings.TrimSpace(out)
	var branch string

	if status != "HEAD" {
		branch = status
	} else {
		branch = getGitDetachedBranch(git, p)
	}

	return []pwl.Segment{{
		Name:       "git-branch",
		Content:    branch,
		Foreground: p.theme.RepoCleanFg,
		Background: p.theme.RepoCleanBg,
	}}
}
