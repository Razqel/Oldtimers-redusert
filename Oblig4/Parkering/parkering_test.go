package Oblig4

import "testing"

func Test_hentOgUnmarshal(t *testing.T) {
	r, _ := hentOgUnmarshal("https://opencom.no/dataset/36ceda99-bbc3-4909-bc52-b05a6d634b3f/resource/d1bdc6eb-9b49-4f24-89c2-ab9f5ce2acce/download/parking.json", &park)

	if r.Sted != "Jernbanen" {
		t.Errorf("return[%v]", r.Sted)
	}

	if r.Dato != "11.05.2018" { //Dato må forandres etter dato den testes.
		t.Errorf("return[%v]", r.Dato)
	}
}