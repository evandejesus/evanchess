package piece

import (
	"testing"
)

func TestIsColor(t *testing.T) {
	type args struct {
		square int
		color  int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "white pawn",
			args: args{
				square: 10,
				color:  White,
			},
			want: true,
		},
		{
			name: "white king",
			args: args{
				square: 9,
				color:  White,
			},
			want: true,
		},
		{
			name: "white queen",
			args: args{
				square: 14,
				color:  White,
			},
			want: true,
		},
		{
			name: "white bishop",
			args: args{
				square: 12,
				color:  White,
			},
			want: true,
		},
		{
			name: "white rook",
			args: args{
				square: 13,
				color:  White,
			},
			want: true,
		},
		{
			name: "white knight",
			args: args{
				square: 11,
				color:  White,
			},
			want: true,
		},
		{
			name: "black pawn",
			args: args{
				square: 18,
				color:  Black,
			},
			want: true,
		},
		{
			name: "black king",
			args: args{
				square: 17,
				color:  Black,
			},
			want: true,
		},
		{
			name: "black queen",
			args: args{
				square: 22,
				color:  Black,
			},
			want: true,
		},
		{
			name: "black bishop",
			args: args{
				square: 20,
				color:  Black,
			},
			want: true,
		},
		{
			name: "black rook",
			args: args{
				square: 21,
				color:  Black,
			},
			want: true,
		},
		{
			name: "black knight",
			args: args{
				square: 19,
				color:  Black,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsColor(tt.args.square, tt.args.color); got != tt.want {
				t.Errorf("IsColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetPieceType(t *testing.T) {
	type args struct {
		square int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "black rook",
			args: args{
				square: 21,
			},
			want: Rook,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPieceType(tt.args.square); got != tt.want {
				t.Errorf("GetPieceType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSlidingPiece(t *testing.T) {
	type args struct {
		square int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "black rook",
			args: args{
				square: 21,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSlidingPiece(tt.args.square); got != tt.want {
				t.Errorf("IsSlidingPiece() = %v, want %v", got, tt.want)
			}
		})
	}
}
