# OSCP Learning Materials - Organized

This directory contains organized OSCP course materials, sanitized for public distribution.

## Directory Structure

```
OSCP-Organized/
├── 1-Basics/                    # Foundation concepts and basics
├── 2-Web-Security/              # Web application security topics
├── 3-Network-Security/          # Network-level vulnerabilities and exploits
├── 4-Exploitation/              # Exploitation techniques and CVEs
├── 5-Post-Exploitation/         # Privilege escalation and post-exploitation
└── 6-Tools-Scripts/             # Reusable scripts and tools
```

## Quick Access via CLI

### Navigate to categories
```bash
# Navigate to Web Security materials
cd 2-Web-Security

# Navigate to Tools
cd 6-Tools-Scripts

# Navigate to Network Security
cd 3-Network-Security

# Navigate to Exploitation
cd 4-Exploitation

# Navigate to Post-Exploitation
cd 5-Post-Exploitation

# Navigate to Basics
cd 1-Basics
```

### View all files in a category
```bash
# List all files in current directory
ls -la
Get-ChildItem  # PowerShell equivalent

# List all markdown files
ls *.md
Get-ChildItem -Filter "*.md"  # PowerShell
```

### Search across all materials
```bash
# Search for specific topics
grep -r "SQL Injection" .
Select-String -Path "*.md" -Pattern "SQL Injection"  # PowerShell

# Search for exploit techniques
grep -r "exploitation" .
grep -r "CVE-" .
```

### View file contents
```bash
# View a markdown file
cat filename.md
Get-Content filename.md  # PowerShell

# View with paging
less filename.md
Get-Content filename.md | more  # PowerShell
```

### Python Scripts
```bash
# Run brute force script
python3 bruteforce.py

# Run with custom target and username
python3 bruteforce.py --target http://TARGET_IP --username admin

# View script source
cat bruteforce.py
```

### Go Scripts
```bash
# Compile Go brute force script
go build -o bruteforce bruteforce.go

# Run compiled binary
./bruteforce

# View Go source
cat bruteforce.go
```

## Content Categories

### 1-Basics
- DNS fundamentals
- Directory listing concepts
- Robots.txt overview
- JavaScript login bypass methods
- Autocomplete vulnerabilities
- Backup lab scenarios

### 2-Web-Security
- SQL Injection techniques (UNION-based, SQLMap POST data)
- Command Injection (with/without filters, blind injection)
- WordPress vulnerabilities
- HTTP Form brute forcing
- Web application reconnaissance

### 3-Network-Security
- SSH brute force techniques
- Tomcat LFI vulnerabilities
- Redis exploitation (Master-Slave mode)
- SMB/Samba vulnerabilities
- Network scanning and reconnaissance

### 4-Exploitation
- CVE-2012-2122: MySQL login bypass
- CVE-2016-10033: WordPress object injection
- CVE-2017-7494: Samba CVE (SambaCry)
- CVE-2017-11610: Supervisord RCE
- CVE-2017-12615: Tomcat PUT method
- CVE-2017-17405: Application-specific RCE
- CVE-2019-5475: Nexus Manager RCE
- CVE-2019-10149: Exim4 RCE
- Metasploit framework usage examples

### 5-Post-Exploitation
- Privilege escalation techniques
- SetUID exploitation
- Crontab abuse
- Race condition exploitation
- System enumeration

### 6-Tools-Scripts
- **bruteforce.py**: HTTP form brute force (Python version)
- **bruteforce.go**: HTTP form brute force (Go version)
- Password wordlists

## Important Notes

⚠️ **Sanitized Content**
- IP addresses have been replaced with `[IP_ADDRESS]` or similar placeholders
- Sensitive credentials have been removed
- Email addresses and personal identifiers have been redacted
- Update configuration variables before using scripts

## Usage Examples

### Find all exploitation techniques
```bash
grep -r "exploit" . --include="*.md"
```

### List all CVE references
```bash
grep -r "CVE-" . --include="*.md"
```

### Search for specific vulnerability types
```bash
grep -r "SQL Injection\|Command Injection\|RCE" . --include="*.md"
```

### View structured index
```bash
find . -type f -name "*.md" | sort
```

## Prerequisites

- Python 3.x (for Python scripts)
- Go 1.15+ (for Go scripts)
- Common penetration testing tools (when applying techniques)
- Password wordlists (provided in some directories)

## Disclaimer

These materials are for educational purposes only. Ensure you have proper authorization before testing any systems. Unauthorized access to computer systems is illegal.

## Last Updated
December 2025
