package domain

type GitHubRepositorySettings struct {
	HasIssues                 bool   `json:"has_issues"`
	HasProjects               bool   `json:"has_projects"`
	HasWiki                   bool   `json:"has_wiki"`
	AllowSquashMerge          bool   `json:"allow_squash_merge"`
	AllowMergeCommit          bool   `json:"allow_merge_commit"`
	AllowRebaseMerge          bool   `json:"allow_rebase_merge"`
	AllowAutoMerge            bool   `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool   `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool   `json:"allow_update_branch"`
	UseSquashPRTitleAsDefault bool   `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitTitle    string `json:"squash_merge_commit_title"`
	SquashMergeCommitMessage  string `json:"squash_merge_commit_message"`
	MergeCommitTitle          string `json:"merge_commit_title"`
	MergeCommitMessage        string `json:"merge_commit_message"`
	Archived                  bool   `json:"archived"`
	WebCommitSignoffRequired  bool   `json:"web_commit_signoff_required"`
}

func NewDefaultGitHubRepositorySettings() *GitHubRepositorySettings {
	return &GitHubRepositorySettings{
		HasIssues:                 true,
		HasProjects:               true,
		HasWiki:                   true,
		AllowSquashMerge:          true,
		AllowMergeCommit:          true,
		AllowRebaseMerge:          true,
		AllowAutoMerge:            false,
		DeleteBranchOnMerge:       false,
		AllowUpdateBranch:         false,
		UseSquashPRTitleAsDefault: false,
		SquashMergeCommitTitle:    "PR_TITLE",
		SquashMergeCommitMessage:  "PR_BODY",
		MergeCommitTitle:          "PR_TITLE",
		MergeCommitMessage:        "PR_BODY",
		Archived:                  false,
		WebCommitSignoffRequired:  false,
	}
}
