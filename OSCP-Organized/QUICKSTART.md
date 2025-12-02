# Quick Start Guide

Get started with OSCP materials in 5 minutes.

## 🚀 5-Minute Quick Start

### Step 1: Navigate to Materials (30 seconds)
```bash
cd OSCP-Organized
```

### Step 2: View Available Categories (30 seconds)
```bash
# Linux/Mac
ls -la
tree -L 1

# PowerShell (Windows)
Get-ChildItem -Directory
```

**Available Categories**:
- 1-Basics
- 2-Web-Security
- 3-Network-Security
- 4-Exploitation
- 5-Post-Exploitation
- 6-Tools-Scripts

### Step 3: Read Main Documentation (2 minutes)
```bash
cat README.md
```

### Step 4: Pick a Category and Dive In (1 minute)
```bash
# Choose one:
cat 1-Basics/README.md
cat 2-Web-Security/README.md
cat 3-Network-Security/README.md
cat 4-Exploitation/README.md
cat 5-Post-Exploitation/README.md
cat 6-Tools-Scripts/README.md
```

### Step 5: Search for Specific Topics (1 minute)
```bash
grep -r "SQL Injection" . --include="*.md"
grep -r "privilege escalation" . --include="*.md"
grep -r "CVE-2017" . --include="*.md"
```

---

## By Experience Level

### 🟢 Beginner
**Start here**:
```bash
# 1. Read the basics
cat 1-Basics/README.md

# 2. Learn web security fundamentals
cat 2-Web-Security/README.md

# 3. Understand the structure
cat INDEX.md
```

**Time**: 1-2 hours

### 🟡 Intermediate
**Focus on**:
```bash
# 1. Network security
cat 3-Network-Security/README.md

# 2. Exploitation techniques
cat 4-Exploitation/README.md

# 3. Practice with tools
cd 6-Tools-Scripts
python3 bruteforce.py
```

**Time**: 2-4 hours

### 🔴 Advanced
**Deep dive into**:
```bash
# 1. Post-exploitation
cat 5-Post-Exploitation/README.md

# 2. Customize and create tools
cd 6-Tools-Scripts
go build bruteforce.go
./bruteforce

# 3. Search for specific CVEs
grep -r "CVE-2019" . --include="*.md"
```

**Time**: 3-8 hours

---

## Common Tasks

### Find Information About...

#### SQL Injection
```bash
grep -r "SQL\|sql injection" . --include="*.md"
cat 2-Web-Security/README.md | grep -A 10 "SQL"
```

#### Command Injection
```bash
grep -r "Command Injection" . --include="*.md"
cat 2-Web-Security/README.md | grep -A 10 "Command"
```

#### Privilege Escalation
```bash
grep -r "privilege\|privilege escalation" . --include="*.md"
cat 5-Post-Exploitation/README.md | head -50
```

#### CVE Exploitation
```bash
# Find all CVEs
grep -r "CVE-" . --include="*.md"

# Find specific CVE
grep -r "CVE-2017-7494" . --include="*.md"

# Find RCE exploits
grep -r "RCE" . --include="*.md"
```

#### Brute Force Attacks
```bash
cat 6-Tools-Scripts/README.md
cd 6-Tools-Scripts
python3 bruteforce.py
```

### Run Tools

#### Python Brute Force
```bash
cd 6-Tools-Scripts
python3 bruteforce.py
# Or configure first:
# - Edit target URL
# - Set username
# - Prepare passwords.txt
```

#### Go Brute Force
```bash
cd 6-Tools-Scripts
go build bruteforce.go
./bruteforce
# Faster than Python version
```

#### Search Across Everything
```bash
grep -r "your_search_term" OSCP-Organized/ --include="*.md"
```

---

## Directory Structure

```
OSCP-Organized/
├── README.md                    ← Start here
├── INDEX.md                     ← Complete index
├── CLI-AUTO-APPROVE-GUIDE.md   ← CLI tips
├── QUICKSTART.md               ← This file
│
├── 1-Basics/                   ← DNS, Directories, Auth
│   └── README.md
├── 2-Web-Security/             ← Injection, Brute Force
│   └── README.md
├── 3-Network-Security/         ← SSH, SMB, Redis, DNS
│   └── README.md
├── 4-Exploitation/             ← CVEs, Metasploit
│   └── README.md
├── 5-Post-Exploitation/        ← Privilege Escalation
│   └── README.md
└── 6-Tools-Scripts/            ← Ready-to-use tools
    ├── README.md
    ├── bruteforce.py
    ├── bruteforce.go
    └── passwords.txt
```

---

## PowerShell Quick Start (Windows)

```powershell
# Navigate
Set-Location OSCP-Organized

# View categories
Get-ChildItem -Directory

# Read README
Get-Content README.md | less

# Search for topics
Select-String -Path "*.md" -Pattern "SQL" -Recurse

# Run Python tool
cd 6-Tools-Scripts
python bruteforce.py
```

---

## Essential Commands Reference

### View Files
```bash
cat filename                # View entire file
head filename              # First 10 lines
tail filename              # Last 10 lines
less filename              # Paginated view
```

### Search
```bash
grep "term" filename       # Search in file
grep -r "term" .          # Search everywhere
grep -i "term" file       # Case-insensitive
grep -v "term" file       # Inverse search
```

### List Files
```bash
ls -la                     # Detailed list
ls -la *.md               # List markdown files
find . -type f -name "*.py"  # Find Python files
```

### Navigation
```bash
cd directory              # Change directory
pwd                       # Current location
cd ..                     # Go back
cd ~                      # Home directory
```

---

## What's in Each Category?

### 1️⃣ Basics (Start here)
- DNS fundamentals
- Directory listing
- Robots.txt
- JavaScript login bypass
- Autocomplete vulnerabilities
- Lab setup

### 2️⃣ Web Security (Most important)
- **SQL Injection**: UNION, time-based, SQLMap
- **Command Injection**: With/without filters, blind
- **WordPress**: Plugin/theme vulnerabilities
- **Brute Force**: HTTP form cracking
- **Tools**: Python & Go scripts included

### 3️⃣ Network Security (Essential)
- SSH brute forcing (Hydra)
- SMB/Samba exploitation
- Redis hacking
- Tomcat vulnerabilities
- DNS enumeration
- Network scanning (nmap)

### 4️⃣ Exploitation (Advanced)
- 8+ CVE exploitation guides
- Metasploit framework
- RCE techniques
- Real-world exploit chains

### 5️⃣ Post-Exploitation (Critical)
- Privilege escalation
- SetUID exploitation
- Crontab abuse
- Race conditions
- Persistence techniques
- Lateral movement

### 6️⃣ Tools & Scripts (Practical)
- Ready-to-use Python tool
- Go version for speed
- Password wordlists
- Complete documentation

---

## Getting Help

### Can't find something?
```bash
# Search globally
grep -r "keyword" . --include="*.md"

# Search by category
grep -r "keyword" 2-Web-Security/ --include="*.md"

# Find files with pattern
find . -name "*sql*" -o -name "*injection*"
```

### Need tool help?
```bash
# View tool README
cat 6-Tools-Scripts/README.md

# Check tool source
cat 6-Tools-Scripts/bruteforce.py
cat 6-Tools-Scripts/bruteforce.go

# View wordlist
cat 6-Tools-Scripts/passwords.txt
```

### Want CLI automation?
```bash
# View full CLI guide
cat CLI-AUTO-APPROVE-GUIDE.md

# View auto-approve examples
cat CLI-AUTO-APPROVE-GUIDE.md | grep -A 20 "Auto-Approve"
```

---

## Recommended Learning Order

**Day 1-2: Basics**
```bash
cat 1-Basics/README.md
cat 2-Web-Security/README.md
```

**Day 3-4: Practice**
```bash
cd 6-Tools-Scripts
python3 bruteforce.py
cat passwords.txt
```

**Day 5-7: Network Security**
```bash
cat 3-Network-Security/README.md
grep -r "SSH\|SMB\|Redis" .
```

**Week 2: Exploitation**
```bash
cat 4-Exploitation/README.md
grep -r "CVE-" .
```

**Week 3: Post-Exploitation**
```bash
cat 5-Post-Exploitation/README.md
grep -r "privilege" .
```

---

## Key Statistics

- **Total categories**: 6
- **Documentation files**: 10+
- **Executable tools**: 2 (Python + Go)
- **Word lists**: 1+
- **CVEs covered**: 8+
- **Topics covered**: 50+
- **Search references**: 100+

---

## Important Notes

⚠️ **Before You Start**:
1. ✅ Read main README.md first
2. ✅ Check INDEX.md for structure
3. ✅ Review CLI guide for execution
4. ✅ Update configuration variables (IPs, usernames)
5. ✅ Only test authorized systems

🔒 **Security Notes**:
- All IPs redacted with [IP_ADDRESS]
- Credentials removed from examples
- For education only
- Obtain authorization before testing

---

## Ready? Let's Go! 🎯

```bash
# Start here
cd OSCP-Organized
cat README.md

# Then pick your path:
cat 1-Basics/README.md
cat 2-Web-Security/README.md
cat 3-Network-Security/README.md
cat 4-Exploitation/README.md
cat 5-Post-Exploitation/README.md
cat 6-Tools-Scripts/README.md

# Happy learning!
```

---

Last Updated: December 2025

For complete information, see: `README.md`, `INDEX.md`, `CLI-AUTO-APPROVE-GUIDE.md`
