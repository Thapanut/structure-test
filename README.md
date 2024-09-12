# structure-test

structure project ที่ใช้จะเป็น pattern Hexagonal Architecture
โดยมี service เป็นตัวหลัก ที่เขียน business logic และจะไปเรียกใช้ชั้นของ repositories ในการ query หรือ เรียกใช้ thirdparty อื่นๆ

# Routes

routes: โฟลเดอร์นี้เก็บไฟล์ที่เกี่ยวข้องกับการกำหนดเส้นทาง (routes) ในแอปพลิเคชัน เช่น การแมป URL ไปยังฟังก์ชันที่เหมาะสม โดยใน routes ประกอบไป ด้วยส่วนต่างๆดังนี้
controllers: โฟลเดอร์นี้เก็บไฟล์ที่เป็นตัวควบคุม (controllers) ซึ่งเป็นส่วนที่รับคำขอจากผู้ใช้ และส่งต่อไปยังส่วนอื่นๆ ของแอปพลิเคชันเพื่อประมวลผลก่อนส่งผลลัพธ์กลับไปยังผู้ใช้
register.go: เป็นไฟล์ Go ที่มีหน้าที่สำคัญในการกำหนดเส้นทาง (path) และวิธีการ (method) ที่แอปพลิเคชันจะรับคำขอ HTTP
search_list.go: เป็น handle ในส่วนของ service searchlist
serverInterface.go: ไฟล์นี้กำหนด interface  สำหรับเซิร์ฟเวอร์ จะทำการ init ก่อนเรียกใช้ ใน register
middleware: โฟลเดอร์นี้เก็บไฟล์ที่เป็น middleware


# Common

ีutils นี้จะเก็บเป็น  funtion ที่เรียกใช้ บ่อยๆ เช่น convertTime
models จะเก็บเป็น struct กลางที่ service เรียกไปใช้ เป็น standard


# stroage

ไว้เกี่ยวกับ ทำ service เชื่อมต่อ stroage เช่น s3 
