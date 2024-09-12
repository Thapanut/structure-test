# Structure-Test

## ภาพรวม
โครงการนี้ใช้ **Hexagonal Architecture** โดยมี **service** เป็นแกนหลักในการประมวลผลธุรกิจ (business logic) และเชื่อมต่อกับส่วนอื่นๆ ผ่าน **repositories**

## โฟลเดอร์หลัก
* **routes:**
  * **controllers:** ประมวลผลคำขอจากผู้ใช้
  * **register.go:** กำหนดเส้นทางและวิธีการรับคำขอ HTTP
  * **search_list.go:** จัดการคำขอค้นหา
  * **serverInterface.go:** กำหนด interface สำหรับเซิร์ฟเวอร์
  * **middleware:** ฟังก์ชันที่ทำงานก่อนและหลังการประมวลผลคำขอ
* **common:**
  * **utils:** ฟังก์ชันทั่วไป เช่น การแปลงเวลา (convertTime)
  * **models:** โครงสร้างข้อมูลที่ใช้ร่วมกัน
* **storage:**
  * จัดการการเชื่อมต่อกับ storage เช่น S3

## Hexagonal Architecture
* **Service:** แกนกลางของระบบ รับผิดชอบในการประมวลผลธุรกิจ
* **Repositories:** ทำหน้าที่เชื่อมต่อกับแหล่งข้อมูล เช่น database, external API

## การใช้งาน
1. **เริ่มต้นเซิร์ฟเวอร์:** เรียกใช้ `main.go`
2. **ส่งคำขอ:** ส่งคำขอ HTTP ไปยังเส้นทางที่กำหนดไว้ใน routes.StartServer() และไปยังไฟล์ `register.go`

## ตัวอย่าง RESTAPI
* **POST /common/na-list/search:** ค้นหาข้อมูล

## อื่นๆ
* **Middleware:** 
  * **Authentication:** ตรวจสอบสิทธิ์ผู้ใช้งาน
  * **Logging:** บันทึก log

## คำถามที่พบบ่อย
* **Q: ** ทำไมต้องใช้ Hexagonal Architecture?
* **A: ** เพื่อแยกส่วนต่างๆ ของระบบออกจากกัน ทำให้ทดสอบและบำรุงรักษาได้ง่ายขึ้น
