# Directory Listing

ให้กระทำการหา folder ลับภายในเว็บไซด์ซึ่งเป็นเป้าหมาย จากนั้นเข้าไปอ่าน flag ที่ถูกซ่อนไว้ใน folder ดังกล่าว

[http://35.240.190.239/dev/secret_YmIxO.txt](http://35.240.190.239/dev/secret_YmIxO.txt)

`dirb <URL>`

`gobuster -w <Wordlist> -u <URL> -t <Thread> -x <EXTENSION>`

`wfuzz -c -w <Wordlist> --hc <ERROR CODE> <URL>/FUZZ`

`dirsearch -u <URL> -e <FILE EXTENSION> -f -x <ERROR CODE>`