package main

import (
	"fmt"
	"net"
)

// กำหนดค่าคงที่สำหรับ username และ password ที่ถูกต้อง
const (
	validUsername = "std1"
	validPassword = "p@ssw0rd"
)

// handleConnection รับเชื่อมต่อและจัดการกับข้อมูลที่ส่งมาจาก Client
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// รับข้อมูลจาก Client
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// แปลงข้อมูลที่รับมาเป็น string
	clientData := string(buffer[:n])

	// ตรวจสอบข้อมูลที่รับมา
	if clientData == fmt.Sprintf("%s:%s", validUsername, validPassword) {
		conn.Write([]byte("Hello\n")) // ถ้าถูกต้อง ส่งข้อความ "Hello" กลับไปที่ Client
	} else {
		conn.Write([]byte("Invalid credentials\n")) // ถ้าไม่ถูกต้อง ส่งข้อความ "Invalid credentials" กลับไปที่ Client
	}
}

func main() {
	fmt.Println("Server is starting...")

	// เปิด port 12345 เพื่อรอรับการเชื่อมต่อ
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()

	for {
		// รอรับการเชื่อมต่อจาก Client
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// เริ่มต้นการจัดการเชื่อมต่อใน Go routine เพื่อรอรับ Client ถัดไป
		go handleConnection(conn)
	}
}
