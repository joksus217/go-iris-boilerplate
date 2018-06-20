package models

type Author struct {
	Id   int64
	Name string
	City string
	Age  int
}

//func (a Author) Validate() (Author, error) {
//	/* do some checks and return an error if that Movie is not valid */
//}
