package articlestore

import (
	"reflect"
	"testing"

	"github.com/GaryBrownEEngr/twertle_api_dev/backend/models"
)

func Test_store_CreateArticle(t *testing.T) {
	type fields struct {
		Articles []models.Article
		NextId   int
	}
	type args struct {
		newArticle *models.Article
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name:   "",
			fields: fields{Articles: []models.Article{}, NextId: 0},
			args:   args{newArticle: &models.Article{Title: "hippos love bacon"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &store{
				Articles: tt.fields.Articles,
				NextId:   tt.fields.NextId,
			}
			a.CreateArticle(tt.args.newArticle)
		})
	}
}

func Test_store_GetArticle(t *testing.T) {
	type fields struct {
		Articles []models.Article
		NextId   int
	}
	type args struct {
		id int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *models.Article
	}{
		// TODO: Add test cases.
		{
			name:   "no return",
			fields: fields{Articles: []models.Article{{Id: 1, Title: "hippos love bacon"}}, NextId: 0},
			args:   args{id: 0},
			want:   nil,
		},
		{
			name:   "no return",
			fields: fields{Articles: []models.Article{{Id: 1, Title: "hippos love bacon"}}, NextId: 0},
			args:   args{id: 1},
			want:   &models.Article{Id: 1, Title: "hippos love bacon"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &store{
				Articles: tt.fields.Articles,
				NextId:   tt.fields.NextId,
			}
			if got := a.GetArticle(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("store.GetArticle() = %v, want %v", got, tt.want)
			}
		})
	}
}
