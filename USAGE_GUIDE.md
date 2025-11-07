# Pentest Assistant - Usage Guide

## Overview
This guide explains how to use the Pentest Assistant AI Agent prompt to capture flags in target systems during OSCP/NCSA exercises.

## Quick Start

### Step 1: Copy the Prompt
Copy the entire content from `PENTEST_ASSISTANT_PROMPT.md` and paste it into your AI assistant (GitHub Copilot, ChatGPT, Claude, etc.)

### Step 2: Provide Target Information
After pasting the prompt, provide your target information in this format:

```
TARGET: [Your target IP/domain]
CATEGORY: [Choose one: Reconnaissance/OSINT | Password Attack | Web Scan | OWASP Top 10]
PROBLEM: [Brief description of what you need to achieve]
HINTS: [Optional: Any hints or clues you have]
```

### Step 3: Let the AI Work
The AI will automatically:
1. Analyze the target and category
2. Select appropriate tools and techniques
3. Execute commands step-by-step
4. Capture and present the flag

## Usage Examples

### Example 1: OSINT on Github
```
TARGET: CompanyXYZ
CATEGORY: Reconnaissance/OSINT
PROBLEM: Find leaked credentials or API keys in Github repositories
HINTS: The company has public repos on Github
```

### Example 2: SSH Brute Force
```
TARGET: 192.168.1.100:22
CATEGORY: Password Attack
PROBLEM: Brute force SSH credentials and retrieve flag.txt
HINTS: Username is likely "admin" with a weak password
```

### Example 3: Web Path Enumeration
```
TARGET: http://webapp.vulnerable-site.com
CATEGORY: Web Scan
PROBLEM: Find hidden directories and sensitive files
HINTS: Look for backup files and configuration files
```

### Example 4: SQL Injection
```
TARGET: http://shop.example.com/product?id=1
CATEGORY: OWASP Top 10
VULNERABILITY: SQL Injection
PROBLEM: Extract flag from database
```

### Example 5: Broken Authentication
```
TARGET: http://portal.example.com/login
CATEGORY: OWASP Top 10
VULNERABILITY: Broken Authentication
PROBLEM: Bypass login and access admin panel
HINTS: Session tokens might be predictable
```

## Supported Categories

### 1. Reconnaissance/OSINT
**Focus:** Information gathering using Google and Github
**Tools:** Google dorks, Github search, trufflehog, git-secrets
**Targets:** Organizations, domains, repositories

### 2. Password Attack
**Focus:** Credential brute forcing with Hydra
**Tools:** Hydra, common wordlists (rockyou.txt)
**Targets:** SSH, FTP, HTTP forms, RDP, MySQL

### 3. Web Scan & Path Enumeration
**Focus:** Discovering hidden web resources
**Tools:** gobuster, dirb, nikto, wpscan, ffuf
**Targets:** Web applications, web servers

### 4. OWASP Top 10
**Focus:** Exploiting common web vulnerabilities
**Tools:** sqlmap, burp suite, custom payloads
**Targets:** Web applications with security flaws

## Tips for Best Results

### 1. Be Specific with Targets
- Provide complete URLs for web targets
- Include port numbers for network services
- Specify exact domains/IPs

### 2. Include Relevant Hints
- Share any reconnaissance findings
- Mention suspected vulnerabilities
- Provide context about the target

### 3. Choose the Right Category
- Match the category to your objective
- Use OWASP Top 10 for web app vulnerabilities
- Use Password Attack specifically for credential brute forcing

### 4. Provide Context
- Share previous findings
- Mention failed attempts
- Include error messages you've seen

## Advanced Usage

### Combining Multiple Attacks
You can chain attacks by providing results from one phase as context for the next:

```
TARGET: 192.168.1.100
CATEGORY: Password Attack
PROBLEM: SSH brute force after OSINT discovered username "dbadmin"
CONTEXT: Previous OSINT found username "dbadmin" in Github repository
```

### Custom Wordlists
Mention if you have custom wordlists:

```
HINTS: Use custom wordlist at /home/user/custom-passwords.txt
```

### Time Constraints
Mention if you need faster or stealthier attacks:

```
HINTS: Target has rate limiting, use slower attack (-t 2 threads)
```

## Tool Prerequisites

Ensure you have these tools installed on your system:

### Basic Tools
```bash
sudo apt update
sudo apt install -y nmap nikto dirb gobuster hydra sqlmap
```

### Additional Tools
```bash
# For OSINT
pip3 install trufflehog
pip3 install gitrob

# For web scanning
go install github.com/ffuf/ffuf@latest
```

### Wordlists
```bash
# Download common wordlists
sudo apt install -y wordlists
# RockYou wordlist location: /usr/share/wordlists/rockyou.txt
# Dirb wordlists: /usr/share/wordlists/dirb/
```

## Output Interpretation

### Success Indicators
Look for these in the AI's response:
- `[+] FLAG CAPTURED:` - Flag successfully found
- `[+] Valid Credentials Found:` - Authentication success
- `[+] Success!` - Exploit worked

### Failure Indicators
- `[-] Failed` - Attempt unsuccessful
- `[*] No results found` - Need different approach
- `[!] Error:` - Technical issue

## Troubleshooting

### Issue: AI suggests tools not installed
**Solution:** Install the suggested tools using apt/pip/go get

### Issue: Commands don't work on your system
**Solution:** Check tool syntax for your OS version, adapt as needed

### Issue: No flag found
**Solution:** Try different category or provide more context/hints

### Issue: Target is unreachable
**Solution:** Verify network connectivity, check firewall rules

## Security & Ethics

### ⚠️ IMPORTANT WARNINGS
1. **Only test systems you own or have explicit permission to test**
2. **Never use these techniques against unauthorized targets**
3. **Respect rate limits and don't perform DoS attacks**
4. **Follow responsible disclosure for real vulnerabilities**

### Legal Considerations
- Unauthorized access is illegal in most jurisdictions
- OSCP labs and CTF environments are authorized testing grounds
- Always get written permission before testing production systems

## Support

### Getting Help
If the AI gets stuck or you need help:
1. Provide more detailed hints
2. Share exact error messages
3. Try a different attack category
4. Manually verify target is reachable

### Community Resources
- OSCP Forums: https://forums.offensive-security.com/
- HackTheBox: https://www.hackthebox.com/
- TryHackMe: https://tryhackme.com/

## Example Full Session

```
=== User Input ===
TARGET: 192.168.1.50
CATEGORY: Password Attack
PROBLEM: Get SSH access and find flag
HINTS: Common username, password in rockyou.txt top 1000

=== AI Response ===
[*] Starting password attack against 192.168.1.50

[Phase 1: Username Enumeration]
Testing common usernames: root, admin, user...

[Phase 2: Hydra Attack]
$ hydra -l admin -P /usr/share/wordlists/rockyou.txt ssh://192.168.1.50 -t 4
[+] Valid credentials: admin:password123

[Phase 3: Access & Flag Retrieval]
$ ssh admin@192.168.1.50
$ find / -name "*flag*" 2>/dev/null
$ cat /home/admin/flag.txt

[+] FLAG CAPTURED: flag{hydra_ssh_pwned}
```

---

## Checklist for Success

Before starting your pentest:
- [ ] Prompt copied to AI assistant
- [ ] Target information gathered
- [ ] Category selected (OSINT/Password/Web/OWASP)
- [ ] Tools installed on local system
- [ ] Network connectivity to target verified
- [ ] Authorization confirmed (own system or CTF/lab)
- [ ] Wordlists available if needed

Happy hacking! 🎯
