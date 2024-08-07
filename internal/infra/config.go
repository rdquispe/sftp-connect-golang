package infra

type Config struct {
	Host           string
	Port           string
	User           string
	PrivateKeyPath string
}

func LoadConfig() *Config {
	return &Config{
		Host:           "<HOST>.server.transfer.us-east-1.amazonaws.com",
		Port:           "<PORT>",
		User:           "<USER_NAME>",
		PrivateKeyPath: "<PATH>/aws_example_ssh_key",
	}
}
