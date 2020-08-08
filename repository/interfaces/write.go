package interfaces

type Write interface {
	Create(t interface{})
	Update(id int64, t interface{})
	Delete(id int64)
}
