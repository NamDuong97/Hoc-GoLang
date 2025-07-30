package main

import (
	"fmt"
)

var (x, y int8)

const MONAN = "CHAO"

// ham trong golang - Hàm có kiểu trả về và trả về 1 biến
// viết hàm trong main và gọi nó trong main thì báo lỗi
func namdaica (age int, name string) string {
	mess := "toi ten la " + name + " nam nay toi da " + fmt.Sprint(age)
	return mess
}

// Hàm trả về 2 biến hoặc nhiều biến, phần trả ve có thể khai báo biến hoặc k cần
func tinhtong (x int, y int) (float32, string) {
	tong:= float32(x) + float32(y)
	mes:= "day la tong cua x va y = " + fmt.Sprint(tong)
	return tong, mes
}

func main() {
	// fmt.Println("nam dai ca")
	// userdaica()
	// // Khai báo biến kiểu var - chỉ khai báo chưa gán giá trị buộc phải có type
	// var nam, tun int 
	// nam = 5
	// tun = 2
	// fmt.Println(nam)
	// fmt.Println(tun)

	// // Khai báo biến kiểu var - chỉ khai báo và có giá trị luôn thì k cần type
	// var nam1, tun1 int  = 12, 15
	// fmt.Println(nam1)
	// fmt.Println(tun1)

	// // Khai báo biến kiểu :=
	// boy, chau := "boy tieu de", "chau heo"
	// fmt.Println(boy)
	// fmt.Println(chau)

	// // Khai báo biến bên hoài hàm main - biến toàn cục sau đó sử dụng trong hàm main
	// x:= 8
	// y:= 10
	// fmt.Println(x)
	// fmt.Println(y)

	// Nhập liệu với scan
	// var hoten string
	// fmt.Println("moi nhap du lieu")
	// // Scan gặp khoảng trắng hay enter nó sẽ tự kết thúc câu lệnh - nếu nhập Trần Doãn Nam thì chỉ nhận dc Trần
	// fmt.Scan(&hoten)
	// fmt.Println("ho ten la: ", hoten)

	// Nhập dữ liệu với bufio - ví dụ nhập tran doan nam thì sẽ lấy cả tran doan nam
	// var hoten1 string
	// fmt.Println("moi nhap du lieu")
	// Scanner := bufio.NewScanner(os.Stdin)
	// if(Scanner.Scan()){
	// 	hoten1 = Scanner.Text()
	// }
	// fmt.Println("ho ten la: ", hoten1)

	// Hằng số giá trị của hằng k dc ghi đè trong cùng 1 phạm vi - block
	// cùng trong hàm main mà có 2 biến const giống nhau sẽ báo lỗi, nhưng ngoài ham main có 1 biến trùng thì vẫn ok, ví dụ biến MONAN
	// const MONAN = "com"
	// const MONAN = "thit" => sẽ báo lỗi vì trong main đã khai báo biến MONAN = "com" rồi nên k dc ghi đè nữa
	// In ra com nếu MONAN trong main còn, nếu MONAN trong main bị comment thì in ra cháo, giá trị biến MONAN ngoai hàm main
	// fmt.Println("mon an la: ", MONAN) 

	// if else
	// diem := 9
	// if diem >= 8 {
	// 	fmt.Println("gioi nhe")
	// } else if diem >= 7 && diem <8 {
	// 	fmt.Println("kha nhe")
	// } else {
	// 	fmt.Println("trung binh nhe")
	// }
	// fmt.Println("done")

	// switch case đối với golang bỏ break cũng k sao cả, chạm phải  đk nào đúng nó sẽ vào chạy xong thì thoát khỏi switch luôn
	// chứ k chạy tiếp tục các câu lệnh khác nữa, case có thể chứa nhiều điều kiện
	// thu := 3
	// switch thu  {
	// 	case 3, 2:
	// 		fmt.Println("thu 3 hoac 2 nhe")
	// 		// break;
	// 	case 4: 
	// 		fmt.Println("thu 4 nhe")
	// 		// break;
	// 	case 5:
	// 		fmt.Println("thu 5 nhe")
	// 		break;
	// 	default: 
	// 		fmt.Println("khong gi ca")
	// 		break;
	// }

	// Vòng lặp for giống for của C++ và C#
	// for i:= 1; i <=10; i++ {
	// 	fmt.Printf("number %d \n", i)
	// }

	// Các biến thể của for và break và continue giống y hệt C++ và C#
	// m:=1 
	// for ; m <=10; {
	// 	fmt.Println("đã vào for") 
	// 	if m%9 == 0 {
	// 		fmt.Printf("number %d \n", m) 
	// 		break
	// 	}
	// 	m++
	// 	if m%5 == 1 {
	// 		continue
	// 	}
	// 	fmt.Printf("number %d \n", m) 
	// 	fmt.Println("đã out for") 
	// }

	// Function trong golang khá tương tự C++ và C# trừ tham số truyền vào thì tên biến trước kiểu dữ liêu
	// fmt.Println(namdaica(28, "nam"))

	// gọi hàm trả về 2 biến
	// a, b:= tinhtong(5, 10)
	// fmt.Println("day la a: ", a)
	// fmt.Println("day la b: ", b)

	// // Ép kiểu trong Golang
	// var a int = 100
	// var s string = fmt.Sprint(a) // int → string dùng fmt.Sprint
	// fmt.Println("day la s ", s)

	// // hoặc dùng strconv
	// s2 := strconv.Itoa(a)       // int → string
	// fmt.Println("day la s2 ", s2)
	// // Ký hiệu _ trong Go được gọi là "blank identifier" 
	// // tức là biến rỗng, dùng để bỏ qua giá trị trả về mà bạn không cần sử dụng.
	// n, _ := strconv.Atoi("123") // string → int
	// fmt.Println("day la n ", n)

	// // Con trỏ là 1 biến để lưu trữ địa chỉ bộ nhớ của biến khác - y chang C++
	// name:= "namdaica"
	// fmt.Println("infomation name variable")
	// fmt.Printf("data type of name: %T \n", name)
	// fmt.Printf("value of name: %v \n", name)
	// fmt.Printf("address of name: %v \n", &name)

	// Tạo con trỏ nè - dùng y chang C++ thôi
	age, old := 8, 10
	var poiter *int 
	poiter = &age
	poiter2 := &old
	fmt.Println("dia chi bien age:", poiter)
	fmt.Println("gia tri bien age:", *poiter)
	fmt.Println("dia chi bien old:", poiter2)
	fmt.Println("gia tri bien old:", *poiter2)
}