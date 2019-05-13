package gitconfig

var defaultConfig = &Config{}

func Do(args ...string) (string, error) {
	return defaultConfig.Do(args...)
}

func Get(args ...string) (string, error) {
	return defaultConfig.Get(args...)
}

func GetAll(args ...string) ([]string, error) {
	return defaultConfig.GetAll(args...)
}

func Bool(key string) (bool, error) {
	return defaultConfig.Bool(key)
}

func Path(key string) (string, error) {
	return defaultConfig.Path(key)
}

func PathAll(key string) ([]string, error) {
	return defaultConfig.PathAll(key)
}

func Int(key string) (int, error) {
	return defaultConfig.Int(key)
}

func User() (string, error) {
	return defaultConfig.User()
}

func Email() (string, error) {
	return defaultConfig.Email()
}

func GitHubToken() (string, error) {
	return defaultConfig.GitHubToken()
}
