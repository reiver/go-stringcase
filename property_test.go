package stringcase


import "testing"


func TestToPropertyCase(t *testing.T) {

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
			expected:  "hello-world",
		},



		{
			thestring: "apple banana cherry",
			expected:  "apple-banana-cherry",
		},
		{
			thestring: "APPLE BANANA CHERRY",
			expected:  "apple-banana-cherry",
		},
		{
			thestring: "Apple Banana Cherry",
			expected:  "apple-banana-cherry",
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
			expected:  "apple-banana-cherry",
		},
		{
			thestring: "APPLE_BANANA_CHERRY",
			expected:  "apple-banana-cherry",
		},
		{
			thestring: "Apple_Banana_Cherry",
			expected:  "apple-banana-cherry",
		},


		{
			thestring: "apple-banana-cherry",
			expected:  "apple-banana-cherry",
		},
		{
			thestring: "APPLE-BANANA-CHERRY",
			expected:  "apple-banana-cherry",
		},
		{
			thestring: "Apple-Banana-Cherry",
			expected:  "apple-banana-cherry",
		},
	}


	for test_number,test := range tests {

		actual := ToPropertyCase(test.thestring)

		if test.expected != actual {
			t.Errorf("For test #%d, for %q expected %q but actually received %q.", test_number, test.thestring, test.expected, actual)
		}
	}
}
