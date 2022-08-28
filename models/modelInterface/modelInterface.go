package modelinterface

type MInterface interface {
	GetOne(id int) MInterface
	GetAll() []MInterface
	Insert(m MInterface)
	Update(m MInterface)
}
