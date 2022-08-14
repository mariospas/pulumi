// nolint: lll
package nodejs

import (
	"testing"
)

func TestMakeSafeEnumName(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
		wantErr  bool
	}{
		{"red", "Red", false},
		{"snake_cased_name", "Snake_cased_name", false},
		{"+", "", true},
		{"*", "Asterisk", false},
		{"0", "Zero", false},
		{"8.3", "TypeName_8_3", false},
		{"11", "TypeName_11", false},
		{"Microsoft-Windows-Shell-Startup", "Microsoft_Windows_Shell_Startup", false},
		{"Microsoft.Batch", "Microsoft_Batch", false},
		{"readonly", "Readonly", false},
		{"SystemAssigned, UserAssigned", "SystemAssigned_UserAssigned", false},
		{"Dev(NoSLA)_Standard_D11_v2", "Dev_NoSLA_Standard_D11_v2", false},
		{"Standard_E8as_v4+1TB_PS", "Standard_E8as_v4_1TB_PS", false},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got, err := makeSafeEnumName(tt.input, "TypeName")
			if (err != nil) != tt.wantErr {
				t.Errorf("makeSafeEnumName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.expected {
				t.Errorf("makeSafeEnumName() got = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestEscape(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input    string
		expected string
	}{
		{"test", "test"},
		{"sub\"string\"", "sub\\\"string\\\""},
		{"slash\\s", "slash\\\\s"},
		{"N\\A \"bad data\"", "N\\\\A \\\"bad data\\\""},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			got := escape(tt.input)
			if tt.expected != got {
				t.Errorf("escape(%s) was %s want %s", tt.input, got, tt.expected)
			}
		})
	}
}
