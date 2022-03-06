package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	return getResult(r, domain)
}

func getResult(r io.Reader, domain string) (DomainStat, error) {
	fileScanner := bufio.NewScanner(r)
	user := new(User)
	result := make(DomainStat)
	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		if err := json.Unmarshal(line, user); err != nil {
			return result, err
		}
		if strings.HasSuffix(user.Email, domain) {
			result[strings.ToLower(strings.Split(user.Email, "@")[1])]++
		}
	}
	return result, nil
}
