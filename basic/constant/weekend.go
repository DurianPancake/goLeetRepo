package constant

type Weekend int

const (
	Monday    Weekend = 1 + iota
	Tuesday           = 1 << iota
	Wednesday         = 1 << iota
	Thursday          = 1 << iota
	Friday            = 1 << iota
	Saturday          = 1 << iota
	Sunday            = 1 << iota
)

type Month int

const (
	January Month = iota + 1
	February
)
