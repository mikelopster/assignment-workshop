
## Assignment Workshop

### workshop-1 = Problem solving

- task-1 = rotation matrix ด้วย python

```
เขียนฟังก์ชัน Python ชื่อ rotate_matrix ที่สามารถหมุนเมทริกซ์ (Matrix) ที่เป็น Input แบบ 2 มิติ (list of lists) ไปทางขวา (ตามเข็มนาฬิกา) 90 องศา
```

```shell
python -m unittest workshop-1/task-1/1_test.py
```

- task-2 = Fizzbuzz ด้วย javascript

```
กฎของ FizzBuzz++:
- ถ้าหารด้วย 3 ลงตัว ให้มีคำว่า "Fizz"
- ถ้าหารด้วย 5 ลงตัว ให้มีคำว่า "Buzz"
- ถ้าหารด้วย 7 ลงตัว ให้มีคำว่า "Woof"
- ถ้าหารด้วยหลายตัวลงตัว ให้นำคำมาต่อกัน (เช่น หารด้วย 3 และ 5 ลงตัว เป็น "FizzBuzz")
- ถ้าหารด้วย 3, 5, หรือ 7 ไม่ลงตัวเลย ให้คืนค่าเป็นตัวเลขนั้นเอง
```

```shell
npm install
npx jest
```

- task-3 = Implement Custom Sorting ด้วย go

```
สร้าง function SortPeopleByAge ที่รับ Slice ของ Person struct ([]Person) และทำการเรียงลำดับ in-place (แก้ไข Slice เดิม) ตามค่า Age จากน้อยไปมาก
```

```shell
go test
```

- task-4 = Binary Tree ด้วย shell script

```
โจทย์: แสดง Binary Tree ด้วย Shell Script (รูปแบบโครงสร้างซ้าย-ขวา)
```

```shell
./task-4/4.sh
```

- workshop-2 = Requirement Analysis
- workshop-3 (try Frontend) = จงทำหน้าเว็บ landing page ให้สวยที่สุด จากเครื่องมืออะไรก็ได้ที่คุณถนัด
- workshop-4 (try Backend) = จงทำ API Support กับโจทย์ UI ดังกล่าว พร้อมทำ Database diagram และ API Specs document (swagger)
- workshop-5 = Testing จงทำ Backend Testing follow ตาม feature และ test case
- workshop-6 = Find Problem in Code