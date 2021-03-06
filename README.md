# gen-interface-mock

It is generator for [gen](https://github.com/cv21/gen) which is able to generate mocks.

Generates [mockery](https://github.com/vektra/mockery)-compatible code for [testify](https://github.com/stretchr/testify).

See examples directory to discover generated code.

### Description

See [DESCRIPTION.md](DESCRIPTION.md) which covers all plugin parameters.

### Example
Input:
```go
package somesvc

import "io"

type SomeService interface {
    SomeSliceGetter() []io.Reader
}
``` 

Result:
```go
// File generated by gen. DO NOT EDIT.
// Generator plugin github.com/cv21/gen-interface-mock v1.0.0.
package bla

import (
	mock "github.com/stretchr/testify/mock"
	"io"
)

// SomeServiceMock is an autogenerated mock type for the SomeService interface.
type SomeServiceMock struct {
	mock.Mock
}

// SomeSliceGetter provides a mock function for method SomeSliceGetter of interface SomeService.
func (_m *SomeServiceMock) SomeSliceGetter() []io.Reader {
	ret := _m.Called()

	var r0 []io.Reader
	if rf, ok := ret.Get(0).(func() []io.Reader); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]io.Reader)
		}
	}

	return r0
}
```

You can see more examples in generator/testdata directory.