// Một	trong	những	kĩ	thuật	quan	trọng	khi	làm	việc	với	database	là	sử	dụng	connection	pool.	Một	connection	xem	như
//  là	một	bộ	đệm	duy	trì	các	kết	nối	tới	database.	Một	connection	pool	là	tập	hợp	nhiều	connection	tới	database.
//  Cơ	chế	hoạt	động	của	connection	pool	khá	đơn	giản,	khi	một	connection	được	tạo	thì	conneciton	đó	sẽ	được	đưa
//  vào	pool	và	được	sử	dụng	lại	cho	các	yêu	cầu	kết	nối	tiếp	theo	cho	đến	khi	bị	đóng	hoặc	hết	thời	gian	chờ	(timeout).
//  Khi	người	dùng	gửi	yêu	cầu	gửi	đến	hệ	thống,	hệ	thống	sẽ	kiểm	tra	xem	trong	pool	có	connection	nào	chưa	đươc	sử
//  dụng	không.	Có	hai	trường	hợp	xảy	ra:
//  Nếu	có	conneciton	chưa	được	sử	dụng,	hệ	thống	sẽ	cung	cấp	connecion	đó	cho	người	dùng	để	xử	lý	các	yêu
//  cầu	kết	nối	tới	database.
//  Nếu	trong	pool	không	rỗng	hoặc	không	có	conneciton	nào	đang	rảnh	và	số	lượng	kết	nối	trong	pool	vẫn	chưa
//  vượt	quá	số	lượng	connection	quy	định	(max	conneciton)	thì	hệ	thống	sẽ	tạo	một	connection	mới	tới	database	và
//  cung	cấp	cho	người	dùng	connection	đó.
//  Nếu	trong	pool	đã	hết	connection	rảnh	và	pool	đã	đạt	số	lượng	conntion	cho	phép	tạo	thì	người	dùng	phải	đợi
//  cho	đến	khi	có	một	connection	rảnh	được	đưa	vào	pool.
//  Sử	dụng	connection	pool	có	nhiều	ưu	điểm:
//  Tăng	hiệu	suất	khi	làm	việc	với	database,	vì	chúng	ta	có	nhiều	kết	nối	tới	database	cùng	lúc	mà	không	phải	đợi
//  tuần	tự.
//  Không	phải	tốn	chi	phí	thời	gian	khởi	tạo	conneciton	và	đóng	connection	cho	mỗi	yêu	cầu	kết	nối	tới	database	vì
//  trong	pool	đã	có	sẵn	connection	được	khởi	tạo	rồi.
//  Sử	dụng	tài	nguyên	hệ	thống	hợp	lý,	khi	chúng	ta	có	thể	tận	dụng	lại	các	connection	đã	sử	dụng	và	giới	hạn
//  được	số	lượng	connection	được	mở.
//  Khi	sử	dụng	package
// database/sql	,	thì	mặc	định	package	đã	hỗ	trợ	chúng	ta	phần	connection	pool.	Nhưng	chúng	ta
//  có	thể	cấu	hình	lại	connection	pool	để	sử	dụng	hiệu	quả	hơn.
//  Sử	dụng	hàm	SetMaxOpenConns	Cấu	hình	số	lượng	connection	lớn	nhất	có	thể	được	mở

package main

// func main() {
// db,	err	:=	sql.Open("mysql",
// "username:password@tcp(127.0.0.1:3306)/hello")
// if	err	!=	nil	{
// log.Fatal(err)
// }
// defer	db.Close()
//	mặc	định	là	0	(không	giới	hạn)
//	nếu	giá	trị	truyền	vào	max	<=	0	cũng	sẽ	là	không	giới	hạn
// db.SetMaxOpenConns(10)
// }

// Sử	dụng	hàm	SetMaxIdleConns	Cấu	hình	số	lượng	connetion	rảnh	có	trong	pool.	Chỉ	số	này	luôn	nhỏ	hơn	hoặc
//  bằng	chỉ	số	MaxOpenConns.	Nếu	chúng	ta	cấu	hình	cao	hơn	thì	thư	việc	sẽ	tự	điều	chỉnh	giảm	lại	cho	phù	hợp	với
//  chỉ	số	MaxOpenConns

// func	main()	{
// 	db,	err	:=	sql.Open("mysql",
// 					"username:password@tcp(127.0.0.1:3306)/hello")
// 	if	err	!=	nil	{
// 					log.Fatal(err)
// 	}
// 	defer	db.Close()
// 	//	mặc	định	là	2	connection	rảnh
// 	//	nếu	giá	trị	truyền	vào	max	<=	0	là	không	có	connection	rãnh	được	giữ	lại
// 	db.SetMaxIdleConns(10)
// }

// 4.5.5	Sử	dụng	Prepared	Statement	để	tăng	hiệu	suất
//  Để	tối	ưu	hiệu	năng	của	hệ	thống,	có	rất	nhiều	cách	để	thực	hiện	nhưng	hiệu	quả	nhất	vẫn	là	tối	ưu	các	câu	truy	vấn
//  database.	Một	trong	số	này	đó	là	sử	dụng		prepared	statement		để	truy	vấn.
//  Prepared	statement	là	một	tính	năng	được	sử	dụng	để	thực	hiện	lặp	lại	các	câu	lệnh	SQL	tương	tự	nhau	với	hiệu	quả
//  cao.	Ví	dụ	minh	hoạ	sau:

// package	main
//  import	(
//  _	"github.com/go-sql-driver/mysql"
//  "database/sql"
//  "log"
//  )
//  func	main(){
//  db,	err	:=	sql.Open("mysql",	"root:root@tcp(127.0.0.1:3306)/hello")
//  if	err	!=	nil	{
//  log.Fatal(err)
//  }
//  stmt,	err	:=	db.Prepare("SELECT	*	FROM	accounts	WHERE	id	=	?;")
//  res,err:=stmt.Exec(2)
//  res,err=stmt.Exec(3)
//  if	err!=nil{
//  log.Fatal(err)
//  }
// log.Println(res)
// }

// Prepare:	đầu	tiên,	ứng	dụng	tạo	ra	1	statement	template	và	gửi	nó	cho	DBMS.	Các	giá	trị	không	được	chỉ	ra	và
//  được	gọi	là
// ?		bên	dưới)
// SELECT	*	FROM	accounts	WHERE	id	=	?;
// Compile:	(parse,	optimizes	và	translates)	statement	template	,	store	kết	quả	mà	không	thực	thi.	Quá	trình	này	do
//  DBMS	thực	hiện.
//  Execute:	ứng	dụng	gửi	giá	trị	của	parametes	của	statement	template	và	DBMS	thực	thi	nó.	Ứng	dụng	có	thể
//  thực	thi	statement	nhiều	lần	với	nhiều	giá	trị	khác	nhau.

// Ưu	điểm	khi	sử	dụng	prepared	statement:
//  Overhead	của	compile	statement	diễn	ra	1	lần	còn	statement	được	thực	thi	nhiều	lần.	Về	lý	thuyết,	khi	sử	dụng
//  prepared	statement,	ta	sẽ	tiết	kiệm	được:
// cost_of_prepare_preprocessing	*	(#statement_executions	-	1)	.	Nhưng
//  thực	tế,	tuỳ	từng	loại	query	sẽ	có	cách	optimize	khác	nhau	(chi	tiết).
//  Chống	SQL	injection.
//  Phát	hiện	sớm	các	lỗi	cú	pháp	trong	câu	statement.
//  Có	thể
// cache	prepared	statement		và	sử	dụng	lại	sau	này.
