package stringcase


import "testing"


func TestToPascalCase(t *testing.T) {

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
			expected:  "HelloWorld",
		},



		{
			thestring: "apple banana cherry",
			expected:  "AppleBananaCherry",
		},
		{
			thestring: "APPLE BANANA CHERRY",
			expected:  "AppleBananaCherry",
		},
		{
			thestring: "Apple Banana Cherry",
			expected:  "AppleBananaCherry",
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
			expected:  "AppleBananaCherry",
		},
		{
			thestring: "APPLE_BANANA_CHERRY",
			expected:  "AppleBananaCherry",
		},
		{
			thestring: "Apple_Banana_Cherry",
			expected:  "AppleBananaCherry",
		},


		{
			thestring: "apple-banana-cherry",
			expected:  "AppleBananaCherry",
		},
		{
			thestring: "APPLE-BANANA-CHERRY",
			expected:  "AppleBananaCherry",
		},
		{
			thestring: "Apple-Banana-Cherry",
			expected:  "AppleBananaCherry",
		},
	}


	for test_number,test := range tests {

		actual := ToPascalCase(test.thestring)

		if test.expected != actual {
			t.Errorf("For test #%d, for %q expected %q but actually received %q.", test_number, test.thestring, test.expected, actual)
		}
	}
}
