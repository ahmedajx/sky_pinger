package env

import "os"

type Environment struct {
	AuthToken  string
	AccountSid string
	From       string
	To         string
}

func Load() Environment {
	env := Environment{
		AuthToken:  os.Getenv("AUTHTOKEN"),
		AccountSid: os.Getenv("ACCOUNTSID"),
		From:       os.Getenv("FROM"),
		To:         os.Getenv("TO"),
	}
	return env
}
