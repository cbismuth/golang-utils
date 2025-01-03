package main

type GitHubRepository struct {
	ID                  int64       `json:"id"`
	NodeID              string      `json:"node_id"`
	Name                string      `json:"name"`
	FullName            string      `json:"full_name"`
	Owner               Owner       `json:"owner"`
	Private             bool        `json:"private"`
	HTMLURL             string      `json:"html_url"`
	Description         string      `json:"description"`
	Fork                bool        `json:"fork"`
	URL                 string      `json:"url"`
	ArchiveURL          string      `json:"archive_url"`
	AssigneesURL        string      `json:"assignees_url"`
	BlobsURL            string      `json:"blobs_url"`
	BranchesURL         string      `json:"branches_url"`
	CollaboratorsURL    string      `json:"collaborators_url"`
	CommentsURL         string      `json:"comments_url"`
	CommitsURL          string      `json:"commits_url"`
	CompareURL          string      `json:"compare_url"`
	ContentsURL         string      `json:"contents_url"`
	ContributorsURL     string      `json:"contributors_url"`
	DeploymentsURL      string      `json:"deployments_url"`
	DownloadsURL        string      `json:"downloads_url"`
	EventsURL           string      `json:"events_url"`
	ForksURL            string      `json:"forks_url"`
	GitCommitsURL       string      `json:"git_commits_url"`
	GitRefsURL          string      `json:"git_refs_url"`
	GitTagsURL          string      `json:"git_tags_url"`
	GitURL              string      `json:"git_url"`
	IssueCommentURL     string      `json:"issue_comment_url"`
	IssueEventsURL      string      `json:"issue_events_url"`
	IssuesURL           string      `json:"issues_url"`
	KeysURL             string      `json:"keys_url"`
	LabelsURL           string      `json:"labels_url"`
	LanguagesURL        string      `json:"languages_url"`
	MergesURL           string      `json:"merges_url"`
	MilestonesURL       string      `json:"milestones_url"`
	NotificationsURL    string      `json:"notifications_url"`
	PullsURL            string      `json:"pulls_url"`
	ReleasesURL         string      `json:"releases_url"`
	SSHURL              string      `json:"ssh_url"`
	StargazersURL       string      `json:"stargazers_url"`
	StatusesURL         string      `json:"statuses_url"`
	SubscribersURL      string      `json:"subscribers_url"`
	SubscriptionURL     string      `json:"subscription_url"`
	TagsURL             string      `json:"tags_url"`
	TeamsURL            string      `json:"teams_url"`
	TreesURL            string      `json:"trees_url"`
	CloneURL            string      `json:"clone_url"`
	MirrorURL           string      `json:"mirror_url"`
	HooksURL            string      `json:"hooks_url"`
	SVNURL              string      `json:"svn_url"`
	Homepage            string      `json:"homepage"`
	Language            interface{} `json:"language"`
	ForksCount          int64       `json:"forks_count"`
	StargazersCount     int64       `json:"stargazers_count"`
	WatchersCount       int64       `json:"watchers_count"`
	Size                int64       `json:"size"`
	DefaultBranch       string      `json:"default_branch"`
	OpenIssuesCount     int64       `json:"open_issues_count"`
	IsTemplate          bool        `json:"is_template"`
	Topics              []string    `json:"topics"`
	HasIssues           bool        `json:"has_issues"`
	HasProjects         bool        `json:"has_projects"`
	HasWiki             bool        `json:"has_wiki"`
	HasPages            bool        `json:"has_pages"`
	HasDownloads        bool        `json:"has_downloads"`
	Archived            bool        `json:"archived"`
	Disabled            bool        `json:"disabled"`
	Visibility          string      `json:"visibility"`
	PushedAt            string      `json:"pushed_at"`
	CreatedAt           string      `json:"created_at"`
	UpdatedAt           string      `json:"updated_at"`
	Permissions         Permissions `json:"permissions"`
	AllowRebaseMerge    bool        `json:"allow_rebase_merge"`
	TemplateRepository  interface{} `json:"template_repository"`
	TempCloneToken      string      `json:"temp_clone_token"`
	AllowSquashMerge    bool        `json:"allow_squash_merge"`
	AllowAutoMerge      bool        `json:"allow_auto_merge"`
	DeleteBranchOnMerge bool        `json:"delete_branch_on_merge"`
	AllowMergeCommit    bool        `json:"allow_merge_commit"`
	SubscribersCount    int64       `json:"subscribers_count"`
	NetworkCount        int64       `json:"network_count"`
	License             License     `json:"license"`
	Forks               int64       `json:"forks"`
	OpenIssues          int64       `json:"open_issues"`
	Watchers            int64       `json:"watchers"`
}

type License struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	SpdxID  string `json:"spdx_id"`
	NodeID  string `json:"node_id"`
	HTMLURL string `json:"html_url"`
}

type Owner struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

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

func NewGitHubRepositorySettings() *GitHubRepositorySettings {
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
