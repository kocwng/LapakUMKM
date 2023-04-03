// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	productTransactions "lapakUmkm/features/productTransactions"

	mock "github.com/stretchr/testify/mock"
)

// ProductTransactionDataInterface is an autogenerated mock type for the ProductTransactionDataInterface type
type ProductTransactionDataInterface struct {
	mock.Mock
}

// Edit provides a mock function with given fields: transactionEntity, id
func (_m *ProductTransactionDataInterface) Edit(transactionEntity productTransactions.ProductTransactionEntity, id uint) error {
	ret := _m.Called(transactionEntity, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(productTransactions.ProductTransactionEntity, uint) error); ok {
		r0 = rf(transactionEntity, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAll provides a mock function with given fields: userId
func (_m *ProductTransactionDataInterface) SelectAll(userId uint) ([]productTransactions.ProductTransactionEntity, error) {
	ret := _m.Called(userId)

	var r0 []productTransactions.ProductTransactionEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]productTransactions.ProductTransactionEntity, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(uint) []productTransactions.ProductTransactionEntity); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]productTransactions.ProductTransactionEntity)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectById provides a mock function with given fields: id
func (_m *ProductTransactionDataInterface) SelectById(id uint) (productTransactions.ProductTransactionEntity, error) {
	ret := _m.Called(id)

	var r0 productTransactions.ProductTransactionEntity
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (productTransactions.ProductTransactionEntity, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uint) productTransactions.ProductTransactionEntity); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(productTransactions.ProductTransactionEntity)
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: transctionEntity
func (_m *ProductTransactionDataInterface) Store(transctionEntity productTransactions.ProductTransactionEntity) (uint, error) {
	ret := _m.Called(transctionEntity)

	var r0 uint
	var r1 error
	if rf, ok := ret.Get(0).(func(productTransactions.ProductTransactionEntity) (uint, error)); ok {
		return rf(transctionEntity)
	}
	if rf, ok := ret.Get(0).(func(productTransactions.ProductTransactionEntity) uint); ok {
		r0 = rf(transctionEntity)
	} else {
		r0 = ret.Get(0).(uint)
	}

	if rf, ok := ret.Get(1).(func(productTransactions.ProductTransactionEntity) error); ok {
		r1 = rf(transctionEntity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductTransactionDataInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductTransactionDataInterface creates a new instance of ProductTransactionDataInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductTransactionDataInterface(t mockConstructorTestingTNewProductTransactionDataInterface) *ProductTransactionDataInterface {
	mock := &ProductTransactionDataInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}