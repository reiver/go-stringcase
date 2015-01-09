package stringcase


import "testing"


func TestToSnakeCase(t *testing.T) {

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
			expected:  "hello_world",
		},



		{
			thestring: "apple banana cherry",
			expected:  "apple_banana_cherry",
		},
		{
			thestring: "APPLE BANANA CHERRY",
			expected:  "apple_banana_cherry",
		},
		{
			thestring: "Apple Banana Cherry",
			expected:  "apple_banana_cherry",
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
			expected:  "apple_banana_cherry",
		},
		{
			thestring: "APPLE_BANANA_CHERRY",
			expected:  "apple_banana_cherry",
		},
		{
			thestring: "Apple_Banana_Cherry",
			expected:  "apple_banana_cherry",
		},


		{
			thestring: "apple-banana-cherry",
			expected:  "apple_banana_cherry",
		},
		{
			thestring: "APPLE-BANANA-CHERRY",
			expected:  "apple_banana_cherry",
		},
		{
			thestring: "Apple-Banana-Cherry",
			expected:  "apple_banana_cherry",
		},
	}


	for test_number,test := range tests {

		actual := ToSnakeCase(test.thestring)

		if test.expected != actual {
			t.Errorf("For test #%d, for %q expected %q but actually received %q.", test_number, test.thestring, test.expected, actual)
		}
	}
}
