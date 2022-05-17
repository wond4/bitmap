package bitmap

import (
	"testing"
)

func TestBitmap_Add(t *testing.T) {
	type fields struct {
		len int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// Add test cases.
		{
			name:   "can add",
			fields: fields{len: 8},
			args:   args{num: 1},
			want:   true,
		}, {
			name:   "can't add",
			fields: fields{len: 8},
			args:   args{num: 10},
			want:   false,
		}, {
			name:   "can't add",
			fields: fields{len: 8},
			args:   args{num: 8},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBitmap(tt.fields.len)
			if got := b.Add(tt.args.num); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitmap_Del(t *testing.T) {
	type fields struct {
		len int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// Add test cases.
		{
			name:   "can del",
			fields: fields{len: 8},
			args:   args{num: 1},
			want:   true,
		}, {
			name:   "can't del",
			fields: fields{len: 8},
			args:   args{num: 10},
			want:   false,
		}, {
			name:   "can't del",
			fields: fields{len: 8},
			args:   args{num: 8},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBitmap(tt.fields.len)
			if got := b.Del(tt.args.num); got != tt.want {
				t.Errorf("Del() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitmap_Has(t *testing.T) {
	type fields struct {
		len int
	}
	type args struct {
		num int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// Add test cases.
		{
			name:   "can has",
			fields: fields{len: 8},
			args:   args{num: 1},
			want:   true,
		}, {
			name:   "can't has",
			fields: fields{len: 8},
			args:   args{num: 10},
			want:   false,
		}, {
			name:   "can't has",
			fields: fields{len: 8},
			args:   args{num: 8},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBitmap(tt.fields.len)
			b.Add(tt.args.num)
			if got := b.Has(tt.args.num); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitmap_Len(t *testing.T) {
	type fields struct {
		len int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// Add test cases.
		{
			name:   "len ok 01",
			fields: fields{len: 8},
			want:   8,
		}, {
			name:   "len ok 02",
			fields: fields{len: 10},
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBitmap(tt.fields.len)
			if got := b.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitmap_LenRaw(t *testing.T) {
	type fields struct {
		len int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// Add test cases.
		{
			name:   "len raw ok 01",
			fields: fields{len: 8},
			want:   1,
		}, {
			name:   "len raw ok 02",
			fields: fields{len: 7},
			want:   1,
		}, {
			name:   "len raw ok 03",
			fields: fields{len: 10},
			want:   2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBitmap(tt.fields.len)
			if got := b.LenRaw(); got != tt.want {
				t.Errorf("LenRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitmap_String(t *testing.T) {
	type fields struct {
		len  int
		adds []int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// Add test cases.
		{
			name:   "string ok 01",
			fields: fields{len: 10},
			want:   "00000000 00",
		}, {
			name:   "string ok 02",
			fields: fields{len: 10, adds: []int{1, 3, 5, 9}},
			want:   "01010100 01",
		}, {
			name:   "string ok 03",
			fields: fields{len: 9, adds: []int{1, 3, 5, 10, 12}},
			want:   "01010100 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBitmap(tt.fields.len)
			for _, i := range tt.fields.adds {
				b.Add(i)
			}
			if got := b.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitmap_index(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// Add test cases.
		{
			name:  "index ok 01",
			args:  args{num: 10},
			want:  1,
			want1: 2,
		}, {
			name:  "index ok 02",
			args:  args{num: 5},
			want:  0,
			want1: 5,
		}, {
			name:  "index ok 03",
			args:  args{num: 200},
			want:  25,
			want1: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bitmap{}
			got, got1 := b.index(tt.args.num)
			if got != tt.want {
				t.Errorf("index() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("index() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		c byte
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// Add test cases.
		{
			name: "reverse ok 01",
			args: args{c: 1},
			want: 128,
		}, {
			name: "reverse ok 02",
			args: args{c: 128},
			want: 1,
		}, {
			name: "reverse ok 03",
			args: args{c: 218},
			want: 91,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.c); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
