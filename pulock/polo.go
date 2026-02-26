package pulock

type Pulock struct {
	Name          string
	Age           int
	Sex           string
	MeritalStatus string
	Phone         string
	Mail          string
	Weight        float32
	Religion      string
}

func NewPulock(name, sex, meritalstatus, phone, mail, religion string, age int, weight float32) *Pulock {
	return &Pulock{Name: name,
		Age:           age,
		Sex:           sex,
		MeritalStatus: meritalstatus,
		Phone:         phone,
		Mail:          mail,
		Weight:        weight,
		Religion:      religion}
}
