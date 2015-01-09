package stringcase


import "testing"


func TestToUpperCase(t *testing.T) {

	tests := []struct{
		thestring string
		expected  string
	}{
		{
			thestring: "test",
			expected:  "TEST",
		},
		{
			thestring: "TEST",
			expected:  "TEST",
		},
		{
			thestring: "Test",
			expected:  "TEST",
		},
		{
			thestring: "tEsT",
			expected:  "TEST",
		},
		{
			thestring: "TeSt",
			expected:  "TEST",
		},



		{
			thestring: "Hello world",
			expected:  "HELLO WORLD",
		},



		{
			thestring: "apple banana cherry",
			expected:  "APPLE BANANA CHERRY",
		},
		{
			thestring: "APPLE BANANA CHERRY",
			expected:  "APPLE BANANA CHERRY",
		},
		{
			thestring: "Apple Banana Cherry",
			expected:  "APPLE BANANA CHERRY",
		},


		{
			thestring: "appleBananaCherry",
			expected:  "APPLEBANANACHERRY",
		},
		{
			thestring: "AppleBananaCherry",
			expected:  "APPLEBANANACHERRY",
		},


		{
			thestring: "apple_banana_cherry",
			expected:  "APPLE_BANANA_CHERRY",
		},
		{
			thestring: "APPLE_BANANA_CHERRY",
			expected:  "APPLE_BANANA_CHERRY",
		},
		{
			thestring: "Apple_Banana_Cherry",
			expected:  "APPLE_BANANA_CHERRY",
		},


		{
			thestring: "apple-banana-cherry",
			expected:  "APPLE-BANANA-CHERRY",
		},
		{
			thestring: "APPLE-BANANA-CHERRY",
			expected:  "APPLE-BANANA-CHERRY",
		},
		{
			thestring: "Apple-Banana-Cherry",
			expected:  "APPLE-BANANA-CHERRY",
		},
	}


	for test_number,test := range tests {

		actual := ToUpperCase(test.thestring)

		if test.expected != actual {
			t.Errorf("For test #%d, for %q expected %q but actually received %q.", test_number, test.thestring, test.expected, actual)
		}
	}
}
