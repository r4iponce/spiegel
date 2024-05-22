package git

type RepoConfig struct {
	URL      string // Source url
	FullPath string // Full clone directory
	Name     string // Name of clone (directory name)
	SSHKey   string // SSH key for auth
}
