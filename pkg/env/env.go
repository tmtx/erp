package env

import (
	"flag"

	"github.com/joho/godotenv"
)

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

var environment = flag.String(
	"env",
	string(Development),
	`Defines current environment.
Possible values are 'development' and 'production'.
Default is 'development'
`,
)

func LoadEnvironment() {
	flag.Parse()

	envFile := ".env"
	if Environment(*environment) == Development {
		envFile = ".env.development"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		panic(err)
	}
}
