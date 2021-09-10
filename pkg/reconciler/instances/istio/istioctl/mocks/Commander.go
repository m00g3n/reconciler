// Code generated by mockery 2.7.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Commander is an autogenerated mock type for the Commander type
type Commander struct {
	mock.Mock
}

// Install provides a mock function with given fields: istioCtlPath, istioOperatorPath, kubeconfigPath
func (_m *Commander) Install(istioCtlPath string, istioOperatorPath string, kubeconfigPath string) error {
	ret := _m.Called(istioCtlPath, istioOperatorPath, kubeconfigPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(istioCtlPath, istioOperatorPath, kubeconfigPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}