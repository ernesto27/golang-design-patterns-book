package visitor

import "testing"

func Test_Overall(t *testing.T) {
	testHelper := &TestHelper{}
	visitor := &MessageVisitor{}

	t.Run("MessageA test", func(t *testing.T) {
		msg := MessageA{
			Msg:    "Hello World",
			Output: testHelper,
		}

		msg.Accept(visitor)
		msg.Print()

		expected := "A: Hello World (Visited A)"
		if testHelper.Received != expected {
			t.Errorf("Expected result was incorrect. %s != %s",
				testHelper.Received, expected)
		}

	})
}
