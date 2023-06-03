package v1

import "testing"

func TestGenerateCodename(t *testing.T) {
	generator := New()

	cases := []struct {
		text   string
		expect string
	}{
		{
			text:   "codename-generator_1",
			expect: "Planet Popsicle",
		},
		{
			text:   "codename-generator_2",
			expect: "Inspiring Sedna",
		},
		{
			text:   "codename-generator_3",
			expect: "Gloomy Churro",
		},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.text, func(t *testing.T) {
			t.Parallel()
			got, err := generator.GenerateCodename(tt.text)
			if err != nil {
				t.Errorf("err %s", err)
			}

			if got != tt.expect {
				t.Errorf("expect = %s, but got = %s", tt.expect, got)
			}
		})
	}
}
