package repository

import (
	"goproject_SBG-bot/datastruct"
	"testing"
)

func TestFind(t *testing.T) {

	//	r := &Repository{}
	//	r.Persons_id = map[int64]*datastruct.Person{}
	//	p:= &datastruct.Person{}
	//	p.ID = 111
	//	p.Date = "200-04-06"
	//	r.Persons_id[p.ID] = p

	//	p := &datastruct.Person{ID: 111, Date: "200-04-06"}
	//	r := &Repository{
	//		Persons_id: map[int64]*datastruct.Person{p.ID: p},
	//	}

	type args struct {
		ID   int64
		Date string
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				ID:   111,
				Date: "200-04-06",
			},
			want: true,
		},
		{
			name: "error1",
			args: args{
				ID:   222,
				Date: "200-04-06",
			},
			want: false,
		},
		{
			name: "error2",
			args: args{
				ID:   111,
				Date: "",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		//		p.ID = tt.args.ID
		//		p.Date = tt.args.Date

		p := &datastruct.Person{ID: tt.args.ID, Date: tt.args.Date}
		r := &Repository{
			Persons_id: map[int64]*datastruct.Person{p.ID: p},
		}

		t.Run(tt.name, func(t *testing.T) {
			got := r.Chek_avtorisation(111)

			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}
}
