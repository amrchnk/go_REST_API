package main

import (
    "fmt"
    "strconv"
    "net/http"
    "github.com/gin-gonic/gin"
)

type ErrorResponse struct{
    Message string `json:"message"`
}

type Handler struct{
    storage Storage
}

func NewHandler(storage Storage) *Handler{
    return &Handler{storage:storage}
}

func (h *Handler) CreateEmployee(c *gin.Context){
    var employee Employee

    if err:=c.BindJSON(&employee);err!=nil{
        fmt.Printf("Failed to bind employee",err.Error())
        c.JSON(http.StatusBadRequest,ErrorResponse{
            Message:err.Error(),
        })
        return
    }

    h.storage.Insert(&employee)

    c.JSON(http.StatusOK,map[string]interface{}{
        "id":employee.ID,
    })
}

func (h *Handler) UpdateEmployee(c *gin.Context){
    id,err:=strconv.Atoi(c.Param("id"))

    if err!=nil{
        fmt.Printf("Failed to convert id param to int: %s\n",err.Error())
        c.JSON(http.StatusBadRequest,ErrorResponse{
                    Message:err.Error(),
        })
        return
    }

    var employee Employee

    if err:=c.BindJSON(&employee);err!=nil{
        fmt.Printf("Failed to bind employee",err.Error())
        c.JSON(http.StatusBadRequest,ErrorResponse{
            Message:err.Error(),
        })
        return
    }

    h.storage.Update(id,employee)

    c.JSON(http.StatusOK,map[string]interface{}{
        "id":employee.ID,
    })
}

func (h *Handler) GetEmployee(c *gin.Context){
    id,err:=strconv.Atoi(c.Param("id"))

    if err!=nil{
            fmt.Printf("Failed to convert id param to int: %s\n",err.Error())
            c.JSON(http.StatusBadRequest,ErrorResponse{
                        Message:err.Error(),
            })
            return
    }

    employee,err:=h.storage.Get(id)

    if err!=nil{
        fmt.Printf("Failed to get employee",err.Error())
                c.JSON(http.StatusBadRequest,ErrorResponse{
                    Message:err.Error(),
                })
        return
    }

    c.JSON(http.StatusOK,employee)
}

func (h *Handler) DeleteEmployee(c *gin.Context){
    id,err:=strconv.Atoi(c.Param("id"))

    if err!=nil{
        fmt.Printf("Failed to convert id param to int: %s\n",err.Error())
        c.JSON(http.StatusBadRequest,ErrorResponse{
            Message:err.Error(),
        })
        return
    }

    h.storage.Delete(id)

    c.String(http.StatusOK,"employee is deleted")
}

func (h *Handler) GetAllEmployees(c *gin.Context){

    employee,err:=h.storage.GetEmployees()

        if err!=nil{
            fmt.Printf("Failed to get employees",err.Error())
                    c.JSON(http.StatusBadRequest,ErrorResponse{
                        Message:err.Error(),
                    })
            return
        }

        c.JSON(http.StatusOK,employee)
}