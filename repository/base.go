package repository

import (
	"github.com/dyshrmn/go-repository-pattern/repository/interfaces"
)

type BaseRepository struct {
	db     map[int64]interface{}
	nextID int64
}

func (b *BaseRepository) Find(id int64) interface{} {
	return b.db[id]
}

func (b *BaseRepository) FindAll() map[int64]interface{} {
	return b.db
}

func (b *BaseRepository) Create(t interface{}) {
	b.db[b.nextID] = t
	b.nextID += 1
}

func (b *BaseRepository) Update(id int64, t interface{}) {
	b.db[b.nextID] = t
}

func (b *BaseRepository) Delete(id int64) {
	delete(b.db, id)
}

func InitBase() *BaseRepository {
	// ensure that this base is implement Read and Write interface
	var _ interfaces.Read = &BaseRepository{}
	var _ interfaces.Write = &BaseRepository{}

	base := BaseRepository{db: make(map[int64]interface{}), nextID: 0}
	return &base
}
