package models

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
	Age  int    `json:"age"`
}

//func (a Author) Validate() (Author, error) {
//	/* do some checks and return an error if that Movie is not valid */
//}
