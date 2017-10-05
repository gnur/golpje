package shows

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/asdine/storm"
	"github.com/gnur/golpje/golpje"
	"github.com/google/uuid"
)

type testDB struct {
	path string
	db   *storm.DB
}

func createEmptyDatabase() testDB {
	u1 := uuid.New()
	randompath := filepath.Join(".", u1.String())
	db, _ := storm.Open(randompath)
	return testDB{
		path: randompath,
		db:   db,
	}
}

func (t testDB) close() {
	t.db.Close()
	os.Remove(t.path)
}

func TestNew(t *testing.T) {
	db := createEmptyDatabase()
	defer db.close()
	type args struct {
		db       *storm.DB
		name     string
		regexp   string
		seasonal bool
		active   bool
		minimal  int64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "add daily show",
			args: args{
				db:       db.db,
				name:     "daily show",
				regexp:   "daily.*show",
				seasonal: false,
				active:   true,
				minimal:  2017,
			},
			wantErr: false,
		},
		{name: "add daily show again",
			args: args{
				db:       db.db,
				name:     "daily show",
				regexp:   "daily.*show",
				seasonal: false,
				active:   true,
				minimal:  2017,
			},
			wantErr: true,
		},
		{name: "empty name fail",
			args: args{
				db:       db.db,
				name:     "",
				regexp:   "daily.*show",
				seasonal: false,
				active:   true,
				minimal:  2017,
			},
			wantErr: true,
		},
		{name: "empty regexp fail",
			args: args{
				db:       db.db,
				name:     "daily show",
				regexp:   "",
				seasonal: false,
				active:   true,
				minimal:  2017,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := New(tt.args.db, tt.args.name, tt.args.regexp, tt.args.seasonal, tt.args.active, tt.args.minimal)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. New() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if _, err := uuid.Parse(got); err != nil && !tt.wantErr {
			t.Errorf("%q. New() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestAll(t *testing.T) {
	db := createEmptyDatabase()
	defer db.close()
	uuid, _ := New(db.db, "daily show", "daily.*show", true, true, 1)
	show := Show{
		ID:       uuid,
		Name:     "daily show",
		Regexp:   "daily.*show",
		Active:   true,
		Seasonal: true,
		Minimal:  1,
	}
	type args struct {
		db *storm.DB
	}
	tests := []struct {
		name    string
		args    args
		want    []Show
		wantErr bool
	}{
		{name: "should get 1 show",
			args: args{
				db: db.db,
			},
			want:    []Show{show},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		got, err := All(tt.args.db)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. All() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. All() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetFromID(t *testing.T) {
	db := createEmptyDatabase()
	defer db.close()
	uuid, _ := New(db.db, "daily show", "daily.*show", true, true, 1)
	show := Show{
		ID:       uuid,
		Name:     "daily show",
		Regexp:   "daily.*show",
		Active:   true,
		Seasonal: true,
		Minimal:  1,
	}
	type args struct {
		db   *storm.DB
		uuid string
	}
	tests := []struct {
		name    string
		args    args
		want    Show
		wantErr bool
	}{
		{name: "get valid uuid",
			args: args{
				db:   db.db,
				uuid: uuid,
			},
			want:    show,
			wantErr: false,
		},
		{name: "get invalid uuid",
			args: args{
				db:   db.db,
				uuid: "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := GetFromID(tt.args.db, tt.args.uuid)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetFromID() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetFromID() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestGetFromName(t *testing.T) {
	db := createEmptyDatabase()
	defer db.close()
	uuid, _ := New(db.db, "daily show", "daily.*show", true, true, 1)
	show := Show{
		ID:       uuid,
		Name:     "daily show",
		Regexp:   "daily.*show",
		Active:   true,
		Seasonal: true,
		Minimal:  1,
	}
	type args struct {
		db   *storm.DB
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    Show
		wantErr bool
	}{
		{name: "get valid name",
			args: args{
				db:   db.db,
				name: "daily show",
			},
			want:    show,
			wantErr: false,
		},
		{name: "get invalid name",
			args: args{
				db:   db.db,
				name: "invalid",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		got, err := GetFromName(tt.args.db, tt.args.name)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GetFromName() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. GetFromName() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestShow_Delete(t *testing.T) {
	db := createEmptyDatabase()
	defer db.close()
	uuid, _ := New(db.db, "daily show", "daily.*show", true, true, 1)
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db *storm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "delete valid show",
			args: args{
				db: db.db,
			},
			fields: fields{
				ID: uuid,
			},
			wantErr: false,
		},
		{name: "delete invalid show",
			args: args{
				db: db.db,
			},
			fields: fields{
				ID: "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		if err := s.Delete(tt.args.db); (err != nil) != tt.wantErr {
			t.Errorf("%q. Show.Delete() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestShow_Print(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db *storm.DB
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		s.Print(tt.args.db)
	}
}

func TestToProtoShows(t *testing.T) {
	type args struct {
		shows []Show
	}
	tests := []struct {
		name string
		args args
		want *golpje.ProtoShows
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := ToProtoShows(tt.args.shows); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. ToProtoShows() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestFromProto(t *testing.T) {
	type args struct {
		in *golpje.ProtoShow
	}
	tests := []struct {
		name string
		args args
		want Show
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		if got := FromProto(tt.args.in); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. FromProto() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestShow_ToProto(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	tests := []struct {
		name   string
		fields fields
		want   *golpje.ProtoShow
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		if got := s.ToProto(); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Show.ToProto() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestShow_ShouldDownload(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db    *storm.DB
		title string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		got, err := s.ShouldDownload(tt.args.db, tt.args.title)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Show.ShouldDownload() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Show.ShouldDownload() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestShow_AddDownload(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db         *storm.DB
		title      string
		magnetlink string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		got, err := s.AddDownload(tt.args.db, tt.args.title, tt.args.magnetlink)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Show.AddDownload() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Show.AddDownload() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestShow_SetDownloaded(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db    *storm.DB
		title string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		if err := s.SetDownloaded(tt.args.db, tt.args.title); (err != nil) != tt.wantErr {
			t.Errorf("%q. Show.SetDownloaded() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestShow_SetDownloadFailed(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db    *storm.DB
		title string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		if err := s.SetDownloadFailed(tt.args.db, tt.args.title); (err != nil) != tt.wantErr {
			t.Errorf("%q. Show.SetDownloadFailed() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestShow_AddEpisode(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db    *storm.DB
		title string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		got, err := s.AddEpisode(tt.args.db, tt.args.title)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Show.AddEpisode() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("%q. Show.AddEpisode() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestShow_DeleteAllEpisodes(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		db *storm.DB
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		if err := s.DeleteAllEpisodes(tt.args.db); (err != nil) != tt.wantErr {
			t.Errorf("%q. Show.DeleteAllEpisodes() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}

func TestShow_Path(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		showBasePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		if got := s.Path(tt.args.showBasePath); got != tt.want {
			t.Errorf("%q. Show.Path() = %v, want %v", tt.name, got, tt.want)
		}
	}
}

func TestShow_GetSeasonDir(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Regexp   string
		Active   bool
		Seasonal bool
		Minimal  int64
	}
	type args struct {
		title        string
		showBasePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		s := Show{
			ID:       tt.fields.ID,
			Name:     tt.fields.Name,
			Regexp:   tt.fields.Regexp,
			Active:   tt.fields.Active,
			Seasonal: tt.fields.Seasonal,
			Minimal:  tt.fields.Minimal,
		}
		if got := s.GetSeasonDir(tt.args.title, tt.args.showBasePath); got != tt.want {
			t.Errorf("%q. Show.GetSeasonDir() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
