package stringcase


import "testing"


func TestToHeaderCase(t *testing.T) {

	tests := []struct{
		thestring string
		expected  string
	}{
		{
			thestring: "test",
			expected:  "Test",
		},
		{
			thestring: "TEST",
			expected:  "Test",
		},
		{
			thestring: "Test",
			expected:  "Test",
		},
		{
			thestring: "tEsT",
			expected:  "Test",
		},
		{
			thestring: "TeSt",
			expected:  "Test",
		},



		{
			thestring: "Hello world",
			expected:  "Hello-World",
		},



		{
			thestring: "apple banana cherry",
			expected:  "Apple-Banana-Cherry",
		},
		{
			thestring: "APPLE BANANA CHERRY",
			expected:  "Apple-Banana-Cherry",
		},
		{
			thestring: "Apple Banana Cherry",
			expected:  "Apple-Banana-Cherry",
		},


		{
			thestring: "appleBananaCherry",
			expected:  "Applebananacherry",
		},
		{
			thestring: "AppleBananaCherry",
			expected:  "Applebananacherry",
		},


		{
			thestring: "apple_banana_cherry",
			expected:  "Apple-Banana-Cherry",
		},
		{
			thestring: "APPLE_BANANA_CHERRY",
			expected:  "Apple-Banana-Cherry",
		},
		{
			thestring: "Apple_Banana_Cherry",
			expected:  "Apple-Banana-Cherry",
		},


		{
			thestring: "apple-banana-cherry",
			expected:  "Apple-Banana-Cherry",
		},
		{
			thestring: "APPLE-BANANA-CHERRY",
			expected:  "Apple-Banana-Cherry",
		},
		{
			thestring: "Apple-Banana-Cherry",
			expected:  "Apple-Banana-Cherry",
		},
	}


	for test_number,test := range tests {

		actual := ToHeaderCase(test.thestring)

		if test.expected != actual {
			t.Errorf("For test #%d, for %q expected %q but actually received %q.", test_number, test.thestring, test.expected, actual)
		}
	}
}
