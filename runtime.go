package requirement

import "runtime"

// ArchitectureIs returns if the current architecture is the same as the specified one
func ArchitectureIs(arch string) func() bool {
	return func() bool {
		return runtime.GOARCH == arch
	}
}

// OperatingSystemIs returns if the current operating system is the same as the specified one
func OperatingSystemIs(os string) func() bool {
	return func() bool {
		return runtime.GOOS == os
	}
}

// GoVersionIs returns if the current go version is the same as the specified one
func GoVersionIs(version string) func() bool {
	return func() bool {
		runtime.Version() == version
	}
}
