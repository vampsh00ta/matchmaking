package models

import "matchmaking/internal/entity"

type Bar struct {
	Age  int
	Name string
}

func (b Bar) encodeToFoo() entity.Foo {
	return entity.Foo{
		Age:  b.Age,
		Name: b.Name,
	}
}
