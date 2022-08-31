package usecase

import "custommer-crud/model"

type CustInteract struct {
	CustRepo CustRepo
}

func (i *CustInteract) Add(c model.Custommer) (cust model.Custommer, err error) {
	cust, err = i.CustRepo.CreateCustomer(c)
	return
}

func (i *CustInteract) CustId(id int) (cust model.Custommer, err error) {
	cust, err = i.CustRepo.GetById(id)
	return
}

func (i *CustInteract) CustAll() (cust model.Custommers, err error) {
	cust, err = i.CustRepo.GetAll()
	return
}

func (i *CustInteract) Update(c model.Custommer) (cust model.Custommer, err error) {
	cust, err = i.CustRepo.UpdateCust(c)
	return
}

func (i *CustInteract) Delete(c model.Custommer) (err error) {
	err = i.CustRepo.DeleteCust(c)
	return
}
