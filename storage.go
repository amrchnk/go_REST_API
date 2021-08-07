package main

import (
    "errors"
    "sync"
)

type Employee struct{
    ID int `json:"id"`
    Name string `json:"name"`
    Sex string `json:"sex"`
    Age int `json:"age"`
    Salary int `json:"salary"`
}

type Employees []Employee

type Storage interface{
    Insert(e *Employee)
    Get(id int)(Employee,error)
    GetEmployees()(Employees,error)
    Update(id int,e Employee)
    Delete(id int)
}

type MemoryStorage struct{
    counter int
    data map[int]Employee
    sync.Mutex
}

func NewMemoryStorage() *MemoryStorage{
    return &MemoryStorage{
        data: make(map[int]Employee),
        counter:1,
    }
}

func (s *MemoryStorage) Insert(e *Employee){
    s.Lock()

    e.ID=s.counter
    s.data[e.ID]=*e

    s.counter++

    s.Unlock()
}

func (s *MemoryStorage) Get(id int)(Employee,error){
    s.Lock()
    defer s.Unlock()

    employee,ok:=s.data[id]
    if !ok{
        return employee,errors.New("Employee isn't found")
    }

    return employee,nil
}

func (s *MemoryStorage) GetEmployees()(Employees,error){
    s.Lock()
    defer s.Unlock()
    var e Employees
    for _,v:=range s.data{
        e=append(e,v)
    }

    return e,nil
}

func (s *MemoryStorage) Update(id int,e Employee){
    s.Lock()
    s.data[id]=e
    s.Unlock()
}

func (s *MemoryStorage) Delete(id int){
    s.Lock()
    delete(s.data,id)
    s.Unlock()
}