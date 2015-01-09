package stringcase


import "testing"


func TestToConstCase(t *testing.T) {

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
			expected:  "HELLO_WORLD",
		},



		{
			thestring: "apple banana cherry",
			expected:  "APPLE_BANANA_CHERRY",
		},
		{
			thestring: "APPLE BANANA CHERRY",
			expected:  "APPLE_BANANA_CHERRY",
		},
		{
			thestring: "Apple Banana Cherry",
			expected:  "APPLE_BANANA_CHERRY",
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
			expected:  "APPLE_BANANA_CHERRY",
		},
		{
			thestring: "APPLE-BANANA-CHERRY",
			expected:  "APPLE_BANANA_CHERRY",
		},
		{
			thestring: "Apple-Banana-Cherry",
			expected:  "APPLE_BANANA_CHERRY",
		},
	}


	for test_number,test := range tests {

		actual := ToConstCase(test.thestring)

		if test.expected != actual {
			t.Errorf("For test #%d, for %q expected %q but actually received %q.", test_number, test.thestring, test.expected, actual)
		}
	}
}
