package hw10programoptimization

import (
	"bufio"
	"fmt"
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
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	for _, user := range u {
		if strings.HasSuffix(user.Email, domain) {
			result[strings.ToLower(strings.Split(user.Email, "@")[1])]++
		}
	}
	return result, nil
}

type users [100_000]User

func getUsers(r io.Reader) (result users, err error) {
	fileScanner := bufio.NewScanner(r)
	num := 0
	user := new(User)
	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		if err = json.Unmarshal(line, user); err != nil {
			return
		}
		result[num] = *user
		num++
	}
	return
}
