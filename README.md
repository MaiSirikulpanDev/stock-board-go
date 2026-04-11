# Stock Board Go

API สำหรับแสดงราคาหุ้นแบบ Real-time โดยใช้เทคโนโลยี WebSocket เพื่อให้ Dashboard สามารถแสดงผลข้อมูลที่อัปเดตทันทีที่ราคาหุ้นเปลี่ยนแปลง

## 🚀 Features

- **WebSocket Support**: รองรับการเชื่อมต่อแบบ Real-time ผ่าน WebSocket
- **Mock Data**: ใช้ข้อมูลจำลอง (Mock Data) สำหรับการทดสอบและพัฒนา
- **Concurrency**: จัดการการเชื่อมต่อหลาย Client พร้อมกันด้วย Goroutines
- **Clean Architecture**: แยกส่วน Repository และ Handler อย่างชัดเจน

## 🛠️ Tech Stack

- **Go 1.25.1**
- **Gin Framework** (สำหรับ HTTP Routing)
- **Gorilla WebSocket** (สำหรับ WebSocket)

## 📂 Project Structure

```
stock-board-go/
├── cmd/
│   └── server/
│       └── main.go          # จุดเริ่มต้นของแอปพลิเคชัน
├── internal/
│   └── stock/
│       ├── controllers/
│       │   └── stock.controller.go # จัดการ HTTP Requests และ WebSocket Connections
│       ├── models/
│       │   └── stock.model.go   # โครงสร้างข้อมูล (Structs)
│       ├── repositories/
│       │   └── stock.repository.go # จัดการการดึงข้อมูล (Mock)
│       └── services/
│           └── stock.service.go # จัดการ Logic ของ API
├── go.mod
└── README.md
```

## ⚙️ Installation

1. **Clone Repository**

   ```bash
   git clone <repository-url>
   cd stock-board-go
   ```

2. **Initialize Go Modules**
   ```bash
   go mod tidy
   ```

## 🏃‍♂️ Running the Server

```bash
go run cmd/server/main.go
```

หรือใช้ `go run` กับไฟล์หลักโดยตรง:

```bash
go run main.go
```

## 🔌 API Endpoints

### 1. HTTP Health Check

ตรวจสอบสถานะของเซิร์ฟเวอร์

- **Method**: `GET`
- **Endpoint**: `/health`
- **Response**: `200 OK`

### 2. WebSocket Connection

เชื่อมต่อเพื่อรับข้อมูลราคาหุ้นแบบ Real-time

- **Method**: `GET`
- **Endpoint**: `/ws/stock`
- **Query Params**: `ticker` (e.g., `?ticker=PTT`)
- **Response**: WebSocket Connection

**ตัวอย่างการเชื่อมต่อ (JavaScript):**

```javascript
const ws = new WebSocket("ws://localhost:8080/ws/stock?ticker=PTT");

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log("Price:", data.price);
  console.log("Last Update:", data.last_update);
};

ws.onclose = () => {
  console.log("Connection closed");
};
```

## 🧪 Testing

### Test Health Check

```bash
curl http://localhost:8080/health
```

### Test WebSocket

ใช้ `wscat` (ติดตั้งผ่าน npm: `npm install -g wscat`) หรือเครื่องมือ WebSocket Client อื่นๆ

```bash
wscat -c ws://localhost:8080/ws/stock?ticker=PTT
```

## 🏗️ Architecture

แอปพลิเคชันนี้ใช้สถาปัตยกรรมแบบ **Clean Architecture** โดยแบ่งเป็น 3 เลเยอร์หลัก:

1. **Repository Layer** (`internal/repositories`)
   - รับผิดชอบการดึงข้อมูล (ปัจจุบันเป็น Mock Data)
   - มี Interface `StockRepository` สำหรับการสลับไปใช้ Database จริงในอนาคต

2. **Handler Layer** (`internal/handlers`)
   - รับ HTTP Requests และ WebSocket Connections
   - ประมวลผล Logic และเรียกใช้ Repository
   - จัดการการส่งข้อมูลผ่าน WebSocket

3. **Model Layer** (`internal/models`)
   - กำหนดโครงสร้างข้อมูล (`structs`) ที่ใช้ในระบบ

## 🔄 Future Enhancements

- [ ] เชื่อมต่อกับ Database จริง (e.g., PostgreSQL, MongoDB)
- [ ] เพิ่มระบบ Authentication และ Authorization
- [ ] รองรับการส่งคำสั่งซื้อขาย (Buy/Sell)
- [ ] เพิ่มกราฟราคา (Historical Data)
- [ ] Implement Rate Limiting
- [ ] เพิ่ม Logging และ Monitoring

## 📝 License

MIT
