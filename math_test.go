package main 
import "testing"

func TestSoma(t *testing.T){
	total:= Sum(10,10)

	if total != 20 {
		t.Errorf("O total esperado era de 20 mas foi %d", total)
	}


}