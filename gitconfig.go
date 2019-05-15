package gitconfig

var defaultConfig = &Config{}

// Do the git config
func Do(args ...string) (string, error) {
	return defaultConfig.Do(args...)
}

// Get a value
func Get(args ...string) (string, error) {
	return defaultConfig.Get(args...)
}

// GetAll values
func GetAll(args ...string) ([]string, error) {
	return defaultConfig.GetAll(args...)
}

// Bool gets a value as bool
func Bool(key string) (bool, error) {
	return defaultConfig.Bool(key)
}

// Path gets a value as path
func Path(key string) (string, error) {
	return defaultConfig.Path(key)
}

// PathAll get all values as paths
func PathAll(key string) ([]string, error) {
	return defaultConfig.PathAll(key)
}

// Int get a value as int
func Int(key string) (int, error) {
	return defaultConfig.Int(key)
}

// User takes git user name
func User() (string, error) {
	return defaultConfig.User()
}

// Email takes git email
func Email() (string, error) {
	return defaultConfig.Email()
}

// GitHubToken takes API token for GitHub
func GitHubToken(host string) (string, error) {
	return defaultConfig.GitHubToken(host)
}

// GitHubUser detects user name of GitHub from various informations
func GitHubUser(host string) (string, error) {
	return defaultConfig.GitHubUser(host)
}
