package restful 



import "testing"


type FromTestValid struct{
	
	name string `required:"true" length:"2,8"`
	email string `required:"true" match:"^[a-z0-9]+@[a-z]+\\.[a-z]{2,3}$"`
}



func TestValidRequired(t *testing.T){
	form := &FromTestValid{email:"test@test.com"}
	
	if _, valid := new(Controller).Valid(form); !valid {
		 t.Errorf("The field required not working with empty field")
	}
}

func TestValidMatch(t *testing.T){
	form := &FromTestValid{name:"test"}
	
	if _, valid := new(Controller).Valid(form); !valid {
		 t.Errorf("The field match not working ")
	}
}

func TestValidLength(t *testing.T){
	form := &FromTestValid{email:"test@test.com",name:"a"}
	
	if _, valid := new(Controller).Valid(form); !valid {
	
		 t.Errorf("length not validation data")
	}
}


func TestValid(t *testing.T){
	form := &FromTestValid{}
	
	if _, valid := new(Controller).Valid(form); !valid {

		 t.Errorf("not validation data with empty struct")
	}
}

