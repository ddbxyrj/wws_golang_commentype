package wwsTranType

type UserLevelMethod interface {
	FindTopInstitutionNum(*[]int) []int
}

type UserLevel struct {
	Org    *[]int
	Posit  *[]int
	Method UserLevelMethod
}
