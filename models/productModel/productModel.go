package productmodel

import m "github.com/VitorC-sa/crud-meplux/models/modelInterface"

type Product struct {
}

func (c Product) GetOne(id int) m.MInterface {
	panic("not implemented") // TODO: Implement
}

func (c Product) GetAll() []m.MInterface {
	panic("not implemented") // TODO: Implement
}

func (c Product) Insert(m m.MInterface) {
	panic("not implemented") // TODO: Implement
}

func (c Product) Update(m m.MInterface) {
	panic("not implemented") // TODO: Implement
}
