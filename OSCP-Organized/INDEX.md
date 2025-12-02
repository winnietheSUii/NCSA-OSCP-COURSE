# OSCP-Organized Index

Comprehensive index of all OSCP learning materials organized by category.

## Main Directories

```
OSCP-Organized/
├── README.md                          # Main guide
├── CLI-AUTO-APPROVE-GUIDE.md         # CLI access & auto-approve instructions
├── INDEX.md                           # This file
├── 1-Basics/
│   └── README.md                      # Basics topics guide
├── 2-Web-Security/
│   └── README.md                      # Web security guide
├── 3-Network-Security/
│   └── README.md                      # Network security guide
├── 4-Exploitation/
│   └── README.md                      # Exploitation guide
├── 5-Post-Exploitation/
│   └── README.md                      # Post-exploitation guide
└── 6-Tools-Scripts/
    ├── README.md                      # Tools & scripts guide
    ├── bruteforce.py                  # Python HTTP brute force
    ├── bruteforce.go                  # Go HTTP brute force
    └── passwords.txt                  # Password wordlist
```

## Quick Access Matrix

| Category | Directory | Primary Topics | Tools |
|----------|-----------|-----------------|-------|
| Basics | 1-Basics | DNS, Directories, Robots.txt, JS Auth | None |
| Web Security | 2-Web-Security | SQL/Command Injection, WordPress, Brute Force | bruteforce.py, bruteforce.go |
| Network Security | 3-Network-Security | SSH, SMB, Redis, Tomcat, DNS | Hydra, nmap, smbclient |
| Exploitation | 4-Exploitation | CVE exploitation, Metasploit | Metasploit, SearchSploit |
| Post-Exploitation | 5-Post-Exploitation | Privilege escalation, persistence | Custom scripts |
| Tools & Scripts | 6-Tools-Scripts | Reusable attack tools | Python, Go, Wordlists |

## Content Overview

### 1-Basics
Essential foundational knowledge:
- DNS fundamentals and enumeration
- Directory listing vulnerabilities
- Robots.txt security implications
- JavaScript-based authentication bypass
- HTML autocomplete exploitation
- Backup and lab setup

**Learning Time**: 2-4 hours

### 2-Web-Security
Web application vulnerabilities:
- SQL Injection (UNION-based, time-based, SQLMap)
- Command Injection (with/without filters, blind)
- WordPress vulnerabilities
- HTTP form brute forcing
- Web reconnaissance techniques

**Tools**: Python/Go brute force scripts, SQLMap, WPScan

**Learning Time**: 5-8 hours

### 3-Network-Security
Network-level security:
- SSH brute forcing (Hydra)
- SMB/Samba vulnerabilities
- Redis exploitation
- Tomcat vulnerabilities
- DNS manipulation
- Network scanning (nmap)

**Tools**: Hydra, nmap, smbclient, redis-cli

**Learning Time**: 4-6 hours

### 4-Exploitation
Vulnerability exploitation:
- CVE-2012-2122: MySQL bypass
- CVE-2016-10033: WordPress/PHPMailer RCE
- CVE-2017-7494: Samba RCE
- CVE-2017-11610: Supervisord RCE
- CVE-2017-12615: Tomcat RCE
- CVE-2019-5475: Nexus RCE
- CVE-2019-10149: Exim4 RCE
- Metasploit framework usage

**Tools**: Metasploit, SearchSploit, Custom exploits

**Learning Time**: 6-10 hours

### 5-Post-Exploitation
Post-compromise actions:
- Linux privilege escalation
- Windows privilege escalation
- SetUID binary exploitation
- Crontab abuse
- Race condition exploitation
- Persistence mechanisms
- Lateral movement
- Data exfiltration

**Tools**: Custom scripts, LinEnum, Privilege Escalation Suggestor

**Learning Time**: 5-8 hours

### 6-Tools-Scripts
Ready-to-use tools:
- **bruteforce.py**: Python HTTP brute force
- **bruteforce.go**: Go HTTP brute force
- **passwords.txt**: Common password wordlist

**Prerequisites**: Python 3.x or Go 1.15+

**Setup Time**: 5-10 minutes

## Access Methods

### Method 1: Direct File Access
```bash
cat OSCP-Organized/README.md
cat OSCP-Organized/1-Basics/README.md
```

### Method 2: Category Navigation
```bash
cd OSCP-Organized/2-Web-Security
ls -la
cat README.md
```

### Method 3: Global Search
```bash
grep -r "SQL Injection" OSCP-Organized/ --include="*.md"
```

### Method 4: Auto-Approve Execution
```bash
source OSCP-Organized/CLI-AUTO-APPROVE-GUIDE.md
# Follow instructions for auto-configuration
```

## Search Index

### By Vulnerability Type

#### Injection Attacks
- SQL Injection: 2-Web-Security/README.md
- Command Injection: 2-Web-Security/README.md
- Blind Injection: 2-Web-Security/README.md

#### Privilege Escalation
- SetUID exploitation: 5-Post-Exploitation/README.md
- Crontab abuse: 5-Post-Exploitation/README.md
- Kernel exploits: 5-Post-Exploitation/README.md
- Sudo misconfiguration: 5-Post-Exploitation/README.md

#### RCE Exploits
- SMB RCE (CVE-2017-7494): 4-Exploitation/README.md
- WordPress RCE (CVE-2016-10033): 4-Exploitation/README.md
- Tomcat RCE (CVE-2017-12615): 4-Exploitation/README.md
- Exim4 RCE (CVE-2019-10149): 4-Exploitation/README.md

#### Authentication Bypass
- MySQL bypass (CVE-2012-2122): 4-Exploitation/README.md
- JavaScript login bypass: 1-Basics/README.md

#### Brute Force Attacks
- HTTP form brute force: 2-Web-Security/README.md
- SSH brute force: 3-Network-Security/README.md

### By Tool

#### Metasploit
- Framework basics: 4-Exploitation/README.md
- Post-exploitation modules: 5-Post-Exploitation/README.md

#### Python
- bruteforce.py: 6-Tools-Scripts/bruteforce.py
- Requirements: requests library

#### Go
- bruteforce.go: 6-Tools-Scripts/bruteforce.go
- No external dependencies

#### Custom Scripts
- SSH Hydra: 3-Network-Security/README.md
- nmap scanning: 3-Network-Security/README.md
- SMB enumeration: 3-Network-Security/README.md

### By CVE

- CVE-2012-2122: MySQL login bypass → 4-Exploitation
- CVE-2016-10033: WordPress/PHPMailer RCE → 4-Exploitation
- CVE-2017-7494: Samba RCE → 4-Exploitation
- CVE-2017-11610: Supervisord RCE → 4-Exploitation
- CVE-2017-12615: Tomcat PUT RCE → 4-Exploitation
- CVE-2017-17405: App RCE → 4-Exploitation
- CVE-2019-5475: Nexus Manager RCE → 4-Exploitation
- CVE-2019-10149: Exim4 RCE → 4-Exploitation

## Quick Reference Commands

### View Content
```bash
cat OSCP-Organized/README.md
less OSCP-Organized/1-Basics/README.md
head OSCP-Organized/2-Web-Security/README.md
```

### Search Topics
```bash
grep -r "CVE-2017" OSCP-Organized/
grep -r "privilege" OSCP-Organized/
grep -r "SQL\|injection" OSCP-Organized/
```

### List Files
```bash
find OSCP-Organized -type f -name "*.md"
find OSCP-Organized -type f -name "*.py"
find OSCP-Organized -type f -name "*.go"
```

### Count Statistics
```bash
find OSCP-Organized -type f | wc -l
grep -r "^##" OSCP-Organized/ | wc -l
wc -l OSCP-Organized/6-Tools-Scripts/passwords.txt
```

## Recommended Learning Sequence

### Phase 1: Foundations (1-2 days)
1. Read: 1-Basics/README.md
2. Review: DNS, Directory listing, Robots.txt
3. Practice: JavaScript auth bypass

### Phase 2: Web Exploitation (2-3 days)
1. Study: 2-Web-Security/README.md
2. Learn: SQL and Command Injection
3. Practice: Use bruteforce tools
4. Exploit: WordPress vulnerabilities

### Phase 3: Network Exploitation (2-3 days)
1. Study: 3-Network-Security/README.md
2. Learn: SSH, SMB, Redis, Tomcat
3. Practice: Network scanning and enumeration
4. Exploit: Service-specific vulnerabilities

### Phase 4: Advanced Exploitation (2-3 days)
1. Study: 4-Exploitation/README.md
2. Learn: CVE exploitation and Metasploit
3. Practice: Exploit database exploration
4. Hands-on: Exploit exercises

### Phase 5: Post-Exploitation (2-3 days)
1. Study: 5-Post-Exploitation/README.md
2. Learn: Privilege escalation
3. Practice: Enumeration scripts
4. Master: Persistence and lateral movement

### Phase 6: Tool Development (1-2 days)
1. Study: 6-Tools-Scripts/README.md
2. Modify: Customize scripts
3. Create: Custom tools
4. Optimize: Performance tuning

**Total Time**: 2-3 weeks for comprehensive coverage

## File Organization Summary

### Documentation Files (7 files)
- README.md
- CLI-AUTO-APPROVE-GUIDE.md
- INDEX.md (this file)
- 1-Basics/README.md
- 2-Web-Security/README.md
- 3-Network-Security/README.md
- 4-Exploitation/README.md
- 5-Post-Exploitation/README.md
- 6-Tools-Scripts/README.md

### Executable Scripts (3 files)
- 6-Tools-Scripts/bruteforce.py
- 6-Tools-Scripts/bruteforce.go
- 6-Tools-Scripts/passwords.txt

### Total Content
- **Directories**: 7
- **Documentation**: 10 files
- **Scripts**: 2 (Python + Go)
- **Wordlists**: 1
- **Total Files**: ~13+

## Notes

⚠️ **Important**
- All IP addresses have been redacted with [IP_ADDRESS] placeholders
- Sensitive credentials have been removed
- This is for educational purposes only
- Always obtain proper authorization before testing
- Follow responsible disclosure practices

## Updates and Modifications

Last Updated: December 2025

**What's New**:
- Organized into 6 distinct categories
- Added CLI and auto-approve guides
- Comprehensive README files for each section
- Executable scripts with configuration templates
- Searchable index for quick access
- Removed critical/sensitive information

**Planned Improvements**:
- Additional automation scripts
- Video resources and links
- Interactive lab environments
- Advanced exploitation techniques
- Certified course materials

## Support Resources

For each topic, refer to:
1. Category-specific README.md
2. CLI-AUTO-APPROVE-GUIDE.md for tool execution
3. Specific vulnerability README.md in subdirectories

Questions? Check:
- Main README.md for structure
- Category README.md for detailed guidance
- CLI guide for execution help
