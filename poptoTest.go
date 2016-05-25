//
// poptoTest.go
// Copyright (C) 2016 Ryan Vlaming <ryanvlaming@icloud.com>
//
// Distributed under terms of the MIT license.
//

package rocket

import "fmt"

func Test() {
	var (
		stack1 Stack
		stack2 Stack
	)

	stack1 = append(stack1, cookie{"some", "thing", 0})
	fmt.Println("stack1 : ", stack1)

	stack1.PopTo(&stack2)

	fmt.Println("stack2 : ", stack2)

}
