package controllers

type Controller interface {
	ReadAllHandler()
	ReadOne(id int64)
	Create()
	Update()
	Delete(id int64)
}
