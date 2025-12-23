# NCSA OSCP Course - Pentest Assistant AI Agent

This repository contains materials for the NCSA OSCP Course Round 1 Post Test, including a comprehensive AI agent prompt designed to assist with penetration testing exercises.

## 🎯 Overview

The **Pentest Assistant AI Agent** is a specially crafted prompt that transforms your AI assistant (GitHub Copilot, ChatGPT, Claude, etc.) into an expert penetration testing companion. It's designed to help you capture flags in target systems across four key OSCP focus areas.

## 📚 Focus Areas

This assistant specializes in:

1. **✓ Reconnaissance / OSINT** (Google & Github) - 1 target
2. **✓ Password Attack using Hydra** - 1 target  
3. **✓ Web Scan & Path Enumeration** - 1 target
4. **✓ OWASP Top 10 Vulnerabilities** - 2 targets

## 🚀 Quick Start

### 1. Get the Prompt
Open [`PENTEST_ASSISTANT_PROMPT.md`](./PENTEST_ASSISTANT_PROMPT.md) and copy the entire content.

### 2. Use with Your AI
Paste the prompt into your preferred AI assistant:
- GitHub Copilot Chat
- ChatGPT
- Claude
- Or any other AI assistant with code execution capabilities

### 3. Provide Target Info
Follow the format:
```
TARGET: [Your target]
CATEGORY: [Reconnaissance/OSINT | Password Attack | Web Scan | OWASP Top 10]
PROBLEM: [What you need to achieve]
HINTS: [Optional clues]
```

### 4. Capture Flags! 🚩
The AI will guide you through the penetration testing process and help you capture flags.

## 📖 Documentation

- **[PENTEST_ASSISTANT_PROMPT.md](./PENTEST_ASSISTANT_PROMPT.md)** - The complete AI agent prompt
- **[USAGE_GUIDE.md](./USAGE_GUIDE.md)** - Detailed usage instructions and examples

## 🎓 Course Materials

The repository also contains practical examples and exercises:

### Network Security
- **SSH Brute Force** - Password attack examples using Hydra
- **Tomcat LFI** - Local File Inclusion exploitation

### Web Security
- **Brute Force HTTP Form Login** - Web authentication attacks
  - Python implementation (`bruteforce.py`)
  - Go implementation (`bruteforce.go`)

### Scanning
- **Metasploit** - Network scanning results and reports

## 🛠️ Required Tools

### Core Tools
```bash
# Install essential pentesting tools
sudo apt update
sudo apt install -y nmap nikto dirb gobuster hydra sqlmap metasploit-framework
```

### OSINT Tools
```bash
# For reconnaissance and OSINT
pip3 install trufflehog gitrob
```

### Web Enumeration
```bash
# Additional web tools
go install github.com/ffuf/ffuf@latest
sudo apt install -y wpscan
```

### Wordlists
```bash
# Common password lists
sudo apt install -y wordlists
# Location: /usr/share/wordlists/rockyou.txt
```

## 💡 Example Usage

### Example 1: SSH Brute Force Attack
```
TARGET: 192.168.1.100:22
CATEGORY: Password Attack
PROBLEM: Brute force SSH and retrieve flag
HINTS: Username is "admin", weak password
```

### Example 2: Web Path Enumeration
```
TARGET: http://vulnerable-app.com
CATEGORY: Web Scan
PROBLEM: Find hidden directories and sensitive files
HINTS: Look for backup and config files
```

### Example 3: SQL Injection
```
TARGET: http://shop.example.com/product?id=1
CATEGORY: OWASP Top 10
VULNERABILITY: SQL Injection
PROBLEM: Extract flag from database
```

## 🎯 Features

The Pentest Assistant AI Agent can:

✅ **Execute any CLI tool** - Full access to system tools  
✅ **Write custom scripts** - Create exploits on the fly  
✅ **Perform reconnaissance** - OSINT via Google and Github  
✅ **Brute force credentials** - Using Hydra and wordlists  
✅ **Enumerate web paths** - Find hidden directories and files  
✅ **Exploit OWASP Top 10** - SQLi, XSS, broken auth, etc.  
✅ **Capture flags** - Primary mission accomplished  
✅ **Document methodology** - Clear step-by-step reporting  

## ⚠️ Legal & Ethical Notice

**IMPORTANT:** This tool is for educational purposes and authorized testing only.

- ✅ **DO USE** on systems you own
- ✅ **DO USE** in CTF competitions
- ✅ **DO USE** in authorized OSCP labs
- ❌ **DO NOT USE** on systems without explicit permission
- ❌ **DO NOT USE** for illegal activities

Unauthorized access to computer systems is illegal. Always obtain written permission before testing any system you don't own.

## 🏆 OSCP Preparation

This repository and the AI assistant are designed to help with:

- **CTF Challenges** - Capture The Flag competitions
- **OSCP Labs** - Offensive Security practice environments
- **Penetration Testing Practice** - Skill development
- **Security Research** - Vulnerability discovery (authorized only)

## 📝 Contributing

This is a course repository. If you have suggestions for improving the AI prompt or adding more examples:

1. Fork the repository
2. Create your feature branch
3. Make your changes
4. Submit a pull request

## 🤝 Support

For questions or issues:
- Review the [Usage Guide](./USAGE_GUIDE.md)
- Check the prompt examples in [PENTEST_ASSISTANT_PROMPT.md](./PENTEST_ASSISTANT_PROMPT.md)
- Consult OSCP community forums

## 📜 License

This repository is for educational purposes as part of the NCSA OSCP Course.

## 🎓 About OSCP

The Offensive Security Certified Professional (OSCP) is a highly respected penetration testing certification. This course material helps prepare students for the practical challenges they'll face in the certification exam.

---

**Remember:** *With great power comes great responsibility. Use these skills ethically and legally.*

Happy hacking! 🚀🔒
