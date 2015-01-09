package stringcase


import "testing"


func TestToCamelCase(t *testing.T) {

	tests := []struct{
		thestring string
		expected  string
	}{
		{
			thestring: "test",
			expected:  "test",
		},
		{
			thestring: "TEST",
			expected:  "test",
		},
		{
			thestring: "Test",
			expected:  "test",
		},
		{
			thestring: "tEsT",
			expected:  "test",
		},
		{
			thestring: "TeSt",
			expected:  "test",
		},



		{
			thestring: "Hello world",
			expected:  "helloWorld",
		},



		{
			thestring: "apple banana cherry",
			expected:  "appleBananaCherry",
		},
		{
			thestring: "APPLE BANANA CHERRY",
			expected:  "appleBananaCherry",
		},
		{
			thestring: "Apple Banana Cherry",
			expected:  "appleBananaCherry",
		},


		{
			thestring: "appleBananaCherry",
			expected:  "applebananacherry",
		},
		{
			thestring: "AppleBananaCherry",
			expected:  "applebananacherry",
		},


		{
			thestring: "apple_banana_cherry",
			expected:  "appleBananaCherry",
		},
		{
			thestring: "APPLE_BANANA_CHERRY",
			expected:  "appleBananaCherry",
		},
		{
			thestring: "Apple_Banana_Cherry",
			expected:  "appleBananaCherry",
		},


		{
			thestring: "apple-banana-cherry",
			expected:  "appleBananaCherry",
		},
		{
			thestring: "APPLE-BANANA-CHERRY",
			expected:  "appleBananaCherry",
		},
		{
			thestring: "Apple-Banana-Cherry",
			expected:  "appleBananaCherry",
		},
	}


	for test_number,test := range tests {

		actual := ToCamelCase(test.thestring)

		if test.expected != actual {
			t.Errorf("For test #%d, for %q expected %q but actually received %q.", test_number, test.thestring, test.expected, actual)
		}
	}
}
