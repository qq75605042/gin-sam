package comm

import "strconv"

type StrTo string

func (f StrTo) MustInt() int {
	v, _ := f.Int()
	return v
}

func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 0)
	return int(v), err
}

func (f StrTo) String() string {
	if f.Exist() {
		return string(f)
	}
}

func (f StrTo) Exist() bool {
	return string(f) != string(0x1E)
}
