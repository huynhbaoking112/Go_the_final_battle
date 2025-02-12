package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// RegisterReq	là	struct	cần	được	validate
type RegisterReq struct {
	//	gt	=	0	cho	biết	độ	dài	chuỗi	phải	>	0，gt:	greater	than
	Username    string `json:"username" validate:"gt=0"`
	PasswordNew string `json:"password_new" validate:"gt=0"`
	//	eqfield	kiểm	tra	các	trường	bằng	nhau
	PasswordRepeat string `json:"password_repeat" validate:"eqfield=PasswordNew"`
	//	kiểm	tra	định	dạng	email	thích	hợp
	Email string `json:"email" validate:"email"`
}

// dùng	1	instance	của	Validate,	cache	lại	struct	info
var validate = validator.New()

// validatefunc	để	wrap	hàm	validate.Struct
func validatefunc(req RegisterReq) error {
	err := validate.Struct(req)
	if err != nil {
		return err
	}
	return nil
}

func main() {

	//	khởi	tạo	obj	để	test	validator
	a := RegisterReq{
		Username:       "Alex",
		PasswordNew:    "",
		PasswordRepeat: "z",
		Email:          "z@z.z",
	}
	err := validatefunc(a)
	fmt.Println(err)
}
