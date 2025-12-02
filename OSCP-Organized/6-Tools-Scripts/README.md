# 6-Tools-Scripts

Reusable penetration testing tools and scripts organized by language and function.

## Available Scripts

### Python Scripts

#### bruteforce.py
HTTP form brute force attack tool

**Purpose**: Brute force HTTP login forms with credential enumeration

**Usage**:
```bash
python3 bruteforce.py
```

**Configuration**:
```python
target = "http://[TARGET_IP]"      # Update with target URL
username = "john"                   # Update with target username
password_file = "passwords.txt"    # Word list file
```

**Features**:
- HTTP POST form submission
- Response analysis for success detection
- Timeout handling
- Adjustable delay between attempts
- Supports multiple success indicators

**Requirements**:
```bash
pip install requests
```

### Go Scripts

#### bruteforce.go
Compiled Go version of HTTP brute force tool

**Purpose**: High-performance HTTP form brute force (compiled binary)

**Compilation**:
```bash
go build -o bruteforce bruteforce.go
```

**Usage**:
```bash
./bruteforce
```

**Configuration**:
```go
target := "http://[TARGET_IP]"     // Update with target URL
username := "john"                  // Update with target username
passwordFile := "passwords.txt"    // Word list file
```

**Features**:
- Concurrent connection handling
- Efficient resource usage
- Redirect following
- Fast execution (compiled binary)
- Response code analysis

**Advantages over Python**:
- Faster execution
- Lower memory footprint
- No runtime required (standalone binary)
- Better for large wordlists

### Word Lists

#### passwords.txt
Common password word list for brute forcing

**Format**: One password per line

**Usage**:
```bash
# View contents
cat passwords.txt

# Count words
wc -l passwords.txt

# View first 10
head -10 passwords.txt

# View last 10
tail -10 passwords.txt

# Use with other tools
hydra -l username -P passwords.txt ssh://[TARGET_IP]
```

## Installation and Setup

### Python Environment

```bash
# Check Python version
python3 --version

# Install dependencies
pip3 install requests

# Run script
python3 bruteforce.py
```

### Go Environment

```bash
# Check Go version
go version

# Build all Go scripts
go build ./...

# Build specific script
go build -o bruteforce bruteforce.go

# Run binary
./bruteforce
```

## Quick Commands

### List all tools
```bash
ls -la
Get-ChildItem  # PowerShell
```

### View Python script
```bash
cat bruteforce.py
less bruteforce.py
```

### View Go script
```bash
cat bruteforce.go
less bruteforce.go
```

### View word list
```bash
cat passwords.txt
less passwords.txt
wc -l passwords.txt  # Count lines
```

### Search for specific functions
```bash
grep -n "def \|class " bruteforce.py  # Python
grep -n "func " bruteforce.go  # Go
```

## Advanced Usage

### Customize brute force script

**Python - Custom target**:
```python
#!/usr/bin/env python3

import requests
from requests.auth import HTTPBasicAuth
import sys
import time

# Your customization here
target = "http://[TARGET_IP]/login.php"
username = "admin"
password_file = "custom_wordlist.txt"
timeout = 5
delay = 0.5
```

**Go - Custom headers**:
```go
// Add custom headers to requests
req, _ := http.NewRequest("POST", target, body)
req.Header.Add("User-Agent", "Custom-Agent")
req.Header.Add("X-Custom-Header", "Value")
```

### Create custom wordlists

```bash
# Generate from common patterns
# Username variations
echo -e "admin\nadministrator\nroot\ntest" > custom.txt

# Merge multiple wordlists
cat wordlist1.txt wordlist2.txt > merged.txt

# Remove duplicates
sort -u merged.txt > merged_unique.txt

# Filter by length
awk 'length > 6' passwords.txt > passwords_min6.txt
```

### Performance tuning

**Python**:
```python
# Reduce delay for faster brute force
time.sleep(0.01)  # 10ms instead of 100ms

# Increase timeout for slow targets
response = requests.post(target, data=data, timeout=30)
```

**Go**:
```go
// Adjust timeout for connections
client := &http.Client{
    Timeout: 30 * time.Second,
}

// Run multiple goroutines for concurrent attempts
go tryLogin(target, username, password)
```

## Security Considerations

⚠️ **Important**:
- Only use on systems with proper authorization
- Be aware of rate limiting and account lockouts
- Some systems log brute force attempts
- Adjust delays to avoid detection
- Test in lab environment first

## Troubleshooting

### Python script issues
```bash
# Import error
pip3 install requests

# Timeout issues
# Edit bruteforce.py, increase timeout value

# Connection refused
# Verify target IP and port
# Check firewall rules
```

### Go script issues
```bash
# Build errors
go get -u ./...

# Runtime errors
# Ensure passwordFile path is correct
# Check file permissions
```

## Customization Examples

### Add logging to Python script
```python
import logging

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

logger.debug(f"Attempting: {username}:{password}")
```

### Add proxy support
```python
proxies = {
    'http': 'http://127.0.0.1:8080',
    'https': 'http://127.0.0.1:8080',
}
response = requests.post(target, data=data, proxies=proxies)
```

### Add threading to Python
```python
from threading import Thread
from queue import Queue

# Create thread pool for parallel requests
threads = 10
q = Queue()

for password in passwords:
    t = Thread(target=try_login, args=(target, username, password))
    t.start()
```

## Quick Reference

```bash
# Directory navigation
pwd
ls -la
cd ../

# File operations
cat filename
head filename
tail filename
wc -l filename
grep pattern filename

# Permissions
chmod +x script.py
chmod +x bruteforce

# Execution
python3 bruteforce.py
./bruteforce
go run bruteforce.go
```

## Prerequisites

- Python 3.6+ (for Python scripts)
- Go 1.15+ (for Go scripts)
- requests library (Python)
- curl/wget (optional, for testing)
- Text editor for script customization

## Additional Resources

- Python Requests Documentation: https://docs.python-requests.org/
- Go HTTP Package: https://golang.org/pkg/net/http/
- HTTP Protocol: https://developer.mozilla.org/en-US/docs/Web/HTTP/
- Authentication Methods: https://owasp.org/www-community/attacks/Brute_force_attack
