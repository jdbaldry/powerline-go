package main

type Symbols struct {
	Lock          string
	Network       string
	Separator     string
	SeparatorThin string

	RepoDetached   string
	RepoAhead      string
	RepoBehind     string
	RepoStaged     string
	RepoNotStaged  string
	RepoUntracked  string
	RepoConflicted string
}

type Theme struct {
	Reset          uint8
	UsernameFg     uint8
	UsernameBg     uint8
	UsernameRootBg uint8

	HostnameFg uint8
	HostnameBg uint8

	HomeSpecialDisplay bool
	HomeFg             uint8
	HomeBg             uint8
	PathFg             uint8
	PathBg             uint8
	CwdFg              uint8
	SeparatorFg        uint8

	ReadonlyFg uint8
	ReadonlyBg uint8

	SshFg uint8
	SshBg uint8

	RepoCleanFg uint8
	RepoCleanBg uint8
	RepoDirtyFg uint8
	RepoDirtyBg uint8

	JobsFg uint8
	JobsBg uint8

	CmdPassedFg uint8
	CmdPassedBg uint8
	CmdFailedFg uint8
	CmdFailedBg uint8

	SvnChangesFg uint8
	SvnChangesBg uint8

	GitAheadFg      uint8
	GitAheadBg      uint8
	GitBehindFg     uint8
	GitBehindBg     uint8
	GitStagedFg     uint8
	GitStagedBg     uint8
	GitNotStagedFg  uint8
	GitNotStagedBg  uint8
	GitUntrackedFg  uint8
	GitUntrackedBg  uint8
	GitConflictedFg uint8
	GitConflictedBg uint8

	VirtualEnvFg uint8
	VirtualEnvBg uint8
}