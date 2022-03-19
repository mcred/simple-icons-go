package simple_icons_go

import (
	"reflect"
	"testing"
)

func TestSimpleIcon_Get(t *testing.T) {
	type args struct {
		slug string
	}
	tests := []struct {
		name    string
		args    args
		want    Icon
		wantErr bool
	}{
		{
			name: "can get by slug",
			args: args{slug: "backbonedotjs"},
			want: Icon{
				Title:   "Backbone.js",
				Slug:    "backbonedotjs",
				Hex:     "0071B5",
				Source:  "https://upload.wikimedia.org/wikipedia/commons/2/20/Backbone.js_logo.svg",
				Svg:     "<svg role=\"img\" viewBox=\"0 0 24 24\" xmlns=\"http://www.w3.org/2000/svg\"><title>Backbone.js</title><path d=\"M2.34 0v10.45l3.2-1.83V5.27l2.93 1.67 3.01-1.72L2.34 0zm19.31 0L12.5 5.22l3.02 1.73 2.94-1.68v3.35l3.2 1.83V0h-.01zm-9.9 5.64-9.4 5.38V24l9.4-5.36v-3.76l-6.21 3.56v-5.5l6.21-3.54V5.64zm.5 0V9.4l6.22 3.54v5.5l-6.22-3.56v3.76L21.66 24V11.02l-9.41-5.38zM7.7 12.3l-1.65.94v1.86l2.17 1.24 3.28-1.87-3.8-2.17zm8.61 0-3.8 2.16 3.28 1.88 2.17-1.24v-1.86l-1.65-.94z\"/></svg>",
				License: License{"MIT", "https://github.com/jashkenas/backbone/blob/master/LICENSE"},
			},
			wantErr: false,
		},
		{
			name:    "expect fail by slug",
			args:    args{slug: "bad_slug_88"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			si := SimpleIcon{}
			got, err := si.Get(tt.args.slug)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}
