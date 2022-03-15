package board

import (
	"reflect"
	"testing"
)

func TestLoadPositionFromFen(t *testing.T) {
	type args struct {
		fen string
	}
	tests := []struct {
		name      string
		args      args
		wantBoard Board
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBoard, err := LoadPositionFromFen(tt.args.fen)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadPositionFromFen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBoard, tt.wantBoard) {
				t.Errorf("LoadPositionFromFen() = %v, want %v", gotBoard, tt.wantBoard)
			}
		})
	}
}
