package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// main เป็นฟังก์ชันหลักที่ทำหน้าที่เป็น Simple Chat Client
func main() {
	fmt.Println("Simple Chat Client")

	// รับป้อน username และ password จากผู้ใช้
	fmt.Print("Connecting to server...\n")
	fmt.Print("Enter username: ")
	username, _ := getUserInput()

	fmt.Print("Enter password: ")
	password, _ := getUserInput()

	// เชื่อมต่อกับ Server ที่รันที่ localhost:12345
	conn, err := net.Dial("tcp", "localhost:12345")
	if handleError(err, "Error connecting to server") {
		return
	}
	defer conn.Close()

	// สร้างข้อมูลที่จะส่งไปยัง Server โดยรวม username และ password
	data := fmt.Sprintf("%s:%s", username, password)

	// ส่งข้อมูลไปยัง Server
	_, err = conn.Write([]byte(data))
	if handleError(err, "Error sending data to server") {
		return
	}

	// รับข้อมูลจาก Server
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if handleError(err, "Error receiving data from server") {
		return
	}

	// แสดงผลลัพธ์ที่ได้จาก Server
	fmt.Println("Server response:", string(buffer[:n]))
}

// getUserInput รับข้อมูลจากผู้ใช้ผ่านทางคีย์บอร์ดและตัดช่องว่างที่เพิ่มเข้ามา
func getUserInput() (string, error) {
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return "", err
	}
	return strings.TrimSpace(input), nil
}

// handleError ตรวจสอบและแสดงข้อผิดพลาด ถ้ามี แล้วคืนค่า true
func handleError(err error, message string) bool {
	if err != nil {
		fmt.Println(message, ":", err)
		return true
	}
	return false
}
