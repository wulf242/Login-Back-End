package authForm

import( 
	"testing"
)
func TestFill(t *testing.T)  {
	Request:= New()
	Request.Fill("me@bsu.edu","kings","0544")
	if Request.email != "me@bsu.edu" || Request.password != "kings"  || Request.token != "0544" {
		t.Error("testFill failed!")
		
	}
	
}


func TestUnpack(t *testing.T)  {
	Request:= New()
	Request.Fill("me@bsu.edu","kings","0544")
	credentials := Request.Unpack()
	if credentials["email"] != "me@bsu.edu" || credentials["password"] != "kings"  || credentials["token"] != "0544" {
		t.Error("testUnpack failed!")	 
	}
	
}