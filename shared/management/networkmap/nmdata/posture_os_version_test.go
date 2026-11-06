package nmdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractNumericVersion(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"6.12.74+deb13+1-amd64", "6.12.74"},
		{"6.12.57+deb13-amd64", "6.12.57"},
		{"6.5.11-linuxkit", "6.5.11"},
		{"6.1.1", "6.1.1"},
		{"10.0.20348.2227", "10.0.20348.2227"},
		{"", ""},
		{"not-a-version", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.want, extractNumericVersion(tt.input))
		})
	}
}

func TestOSVersionCheck_DebianKernel(t *testing.T) {
	check := &OSVersionCheck{
		Linux: &MinKernelVersionCheck{
			MinKernelVersion: "0",
		},
	}

	peer := &Peer{
		Meta: PeerSystemMeta{
			GoOS:          "linux",
			KernelVersion: "6.12.74+deb13+1-amd64",
		},
	}

	isValid, err := check.check(peer)
	assert.NoError(t, err)
	assert.True(t, isValid, "Debian 13 kernel with +deb13+1 suffix should pass min version 0")
}
