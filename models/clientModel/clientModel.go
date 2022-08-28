package clientmodel

import m "github.com/VitorC-sa/crud-meplux/models/modelInterface"

type Client struct {
}

func (c *Client) GetOne(id int) m.MInterface {
	panic("not implemented") // TODO: Implement
}

func (c *Client) GetAll() []m.MInterface {
	panic("not implemented") // TODO: Implement
}

func (c *Client) Insert(m m.MInterface) {
	panic("not implemented") // TODO: Implement
}

func (c *Client) Update(m m.MInterface) {
	panic("not implemented") // TODO: Implement
}
