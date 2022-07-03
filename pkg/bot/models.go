package bot

type User struct {
	FName string
	SName string
	Date  string
	State int // 0 by default, 1 for entering FName, 2 for entering SName, 3 for entering Date 
}
