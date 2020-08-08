package interfaces

type Read interface {
	Find(id int64) interface{}
	FindAll() map[int64]interface{}
}
