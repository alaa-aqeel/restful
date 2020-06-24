package restful 


import "os"
import "testing"



// Test Create Token With out Secret key 
func TestCreateTokenWithoutSecretKey(t *testing.T){

	if _, err := new(Controller).CreateToken("1", 1); err == nil {

		t.Error("With Out secret CreateToken working")
	}
}

// Here test Create Token With Secret key  antd with out id user 
func TestCreateTokenWithIdEmpty(t *testing.T){

	os.Setenv("TOEKN_SECRET_KEY", "TEST_SECRET_KEY")

	if _, err := new(Controller).CreateToken("", 1); err == nil {
		
		t.Error("With empty id get token not nill error ")
	}
}

// Here test Verify Token With valid Token 
func TestVerifyToken(t *testing.T){

	os.Setenv("TOEKN_SECRET_KEY", "TEST_SECRET_KEY")

	if token, err := new(Controller).CreateToken("", 1); err == nil {
		
		if tkn, _,  err := VerifyToken(token); err != nil && !tkn.Valid {

			t.Error("Verify Token Not Work")
		}
	}
}

// Here test Verify Token With invalid Token 
func TestVerifyTokenExpiresToken(t *testing.T){

	// This token is expires
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTI5NDI5MTAsImp0aSI6IjMifQ.PAYEsYB-QQ8hjXFaOwhVkeLrApPkSMRD9Hdn4QP0M3Q"

	if tkn, _,  err := VerifyToken(token); err == nil && tkn.Valid {

		t.Error("Verify Token with expires token not working")
	}
}
