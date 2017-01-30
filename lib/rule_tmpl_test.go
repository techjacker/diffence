package diffence

import "testing"

func Test_rule_Run(t *testing.T) {
	type fields struct {
		Caption     string
		Description interface{}
		Part        string
		Pattern     string
		Type        string
	}
	type args struct {
		in string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Should find a match",
			fields: fields{
				Caption:     "Contains word: password",
				Description: nil,
				Part:        "filename",
				Pattern:     "password",
				Type:        "regex",
			},
			args: args{"/path/to/password.txt"},
			want: true,
		},
		{
			name: "Should not find a match (only check filename not preceding path)",
			fields: fields{
				Caption:     "Contains word: password",
				Description: nil,
				Part:        "filename",
				Pattern:     "password",
				Type:        "regex",
			},
			args: args{"/path/to/password/file.txt"},
			want: false,
		},
		{
			name: "Should find a match",
			fields: fields{
				Caption:     "Potential cryptographic private key",
				Description: nil,
				Part:        "extension",
				Pattern:     "pem",
				Type:        "match",
			},
			args: args{"/path/to/password.pem"},
			want: true,
		},
		{
			name: "Should not find a match (only check extension only not preceding path)",
			fields: fields{
				Caption:     "Potential cryptographic private key",
				Description: nil,
				Part:        "extension",
				Pattern:     "pem",
				Type:        "match",
			},
			args: args{"/path/to/pem.txt"},
			want: false,
		},
		{
			name: "Should find a match in extension",
			fields: fields{
				Caption:     "Ruby On Rails database schema file",
				Description: "Contains information on the database schema of a Ruby On Rails application.",
				Part:        "filename",
				Pattern:     "schema.rb",
				Type:        "match",
			},
			args: args{"/path/to/schema.rb"},
			want: true,
		},
		{
			name: "Should not find a match (only check filename not preceding path)",
			fields: fields{
				Caption:     "Ruby On Rails database schema file",
				Description: "Contains information on the database schema of a Ruby On Rails application.",
				Part:        "filename",
				Pattern:     "schema.rb",
				Type:        "match",
			},
			args: args{"/path/to/schema.rb/different/file.txt"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rule{
				Caption:     tt.fields.Caption,
				Description: tt.fields.Description,
				Part:        tt.fields.Part,
				Pattern:     tt.fields.Pattern,
				Type:        tt.fields.Type,
			}
			if got := r.Run(tt.args.in); got != tt.want {
				t.Errorf("rule.Run() = %v, want %v", got, tt.want)
			}
		})
	}
}
