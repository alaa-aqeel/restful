/*
	Package restful implements a simple library for restful-api with fiber.
	
*/
package restful 

import (
	"reflect"
	"strconv"
	"fmt"
	"strings"
	"regexp"
)

// fileds :
// 	required :
// 		The field under validation must be present in the input data and not empty
// 	length:
// 		The field under validation must have a size between the given min and max
// 	match:
// 		The field under validation must match the given regular expression.

// The funcation Check length between the given min and max. 
func Length(rule string, length int) bool{
	
	split := strings.Split(rule, ",")
	man, _ := strconv.Atoi(split[0])
	max, _ := strconv.Atoi(split[1])

	return (man > length || length > max ) 
}

// match with regular expression.
func Match(pattern string, value string) bool {

	match, _ := regexp.MatchString(pattern, value)
	return match 
}


/* 	
	Valid implements value validations for structs

	parameter:
		models: Fileds struct  

	return map[string]interface{}, bool

*/
func (ctrl *Controller) Valid(models interface{}) (map[string]interface{}, bool) {

	errorMap := make(map[string]interface{})
	model := reflect.ValueOf(models).Elem()

	for i := 0; i < model.NumField(); i++ {
	    field  := model.Type().Field(i)
	    errors := make([]string, 0)
	    errorsMsg, _ := ctrl.MessageValid[field.Name] 


	    if _, ok := field.Tag.Lookup("required"); ok {
		    if  strings.TrimSpace(model.Field(i).String()) == "" {
		    	if msg, ok := errorsMsg["required"]; ok {
		    		errors = append(errors, fmt.Sprintf(msg, field.Name))
		    	} else {
		    	 	errors = append(errors, 
		    	 		fmt.Sprintf("The %s field is required", field.Name))
		    	}
		    }
		}

		if pattern, ok := field.Tag.Lookup("match"); ok {
			
			if !Match(pattern, model.Field(i).String()) {
				if msg, ok:= errorsMsg["match"]; ok{
		    		errors = append(errors, fmt.Sprintf(msg, field.Name))
		    	} else {
		    	 	errors = append(errors, 
		    	 		fmt.Sprintf("The %s field format is invalid.", field.Name))
		    	}
			}
		}

	    if rule , ok := field.Tag.Lookup("length"); ok {
	    	if Length(rule, len(model.Field(i).String()) ){
	    		
		    	if msg, ok:= errorsMsg["length"]; ok{

		    		errors = append(errors,fmt.Sprintf(msg, field.Name, rule))		    	
		    	} else {

		    	 	errors = append(errors, 
		    	 		fmt.Sprintf("The %s field must be between %s", field.Name, rule))
		    	}
	    	}
	    }

	    if len(errors) != 0{

	    	errorMap[field.Name] = errors 
	    }
	    
	}

	return errorMap, len(errorMap) != 0
}