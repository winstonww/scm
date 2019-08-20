package core

import (
	"math"
	"reflect"
)

type Comparable interface {
	LessThan(j interface{}) bool
	GreaterThan(j interface{}) bool
	EqualTo(j interface{}) bool
}

type Addable interface {
	//TODO Need a better name for this interface
	Add(j interface{}) interface{}
	Subtract(j interface{}) interface{}
}

type Load struct {
	/* The load here is 2 dimentional but can be extended
	   to k-dimentional. Members of the struct must be of
	   type float64
	*/
	CPU       float64
	Bandwidth float64
}

func (x Load) EqualTo(y Load) bool {
	xx, yy := reflect.ValueOf(x), reflect.ValueOf(y)
	for i := 0; i < xx.NumField(); i++ {
		if xx.Field(i) != yy.Field(i) {
			return false
		}
	}
	return true
}

func (x Load) LessThan(y Load) bool {
	xx, yy := reflect.ValueOf(x), reflect.ValueOf(y)
	for i := 0; i < xx.NumField(); i++ {
		if xload, ok := xx.Field(i).Interface().(float64); ok {
			if yload, ok := yy.Field(i).Interface().(float64); ok {
				if xload > yload {
					return false
				}
			}
		}
	}
	return true
}

func (x Load) GreaterThan(y Load) bool {
	return !x.LessThan(y) && !x.EqualTo(y)
}

/* return a new Load by adding Load x to Load y */
func (x Load) Add(y Load) Load {
	xx, yy := reflect.ValueOf(x), reflect.ValueOf(y)
	res := Load{}
	resField := reflect.ValueOf(&res).Elem()
	for i := 0; i < resField.NumField(); i++ {
		if xload, ok := xx.Field(i).Interface().(float64); ok {
			if yload, ok := yy.Field(i).Interface().(float64); ok {
				if resField.Field(i).CanSet() {
					resField.Field(i).SetFloat(xload + yload)
				}
			}
		}
	}
	return res
}

/* return a new Load by subracting Load y from Load x */
func (x Load) Subtract(y Load) Load {
	xx, yy := reflect.ValueOf(x), reflect.ValueOf(y)
	res := Load{}
	resField := reflect.ValueOf(&res).Elem()
	for i := 0; i < resField.NumField(); i++ {
		if xload, ok := xx.Field(i).Interface().(float64); ok {
			if yload, ok := yy.Field(i).Interface().(float64); ok {
				resField.Field(i).SetFloat(xload - yload)
			}
		}
	}
	return res
}

func (x Load) L2Norm() float64 {
	norm, xx := 0.0, reflect.ValueOf(x)
	for i := 0; i < xx.NumField(); i++ {
		if xload, ok := xx.Field(i).Interface().(float64); ok {
			norm += xload * xload
		}
	}
	return math.Sqrt(norm)
}
