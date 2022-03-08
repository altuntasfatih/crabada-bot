package storage

import (
	"github.com/patrickmn/go-cache"
	"reflect"
	"testing"
	"time"
)

var (
	getMock   func(key string) (interface{}, bool)
	setMock   func(key string, value interface{}, d time.Duration)
	flushMock func()
	itemsMock func() map[string]cache.Item
)

type setGetFlusherMock struct {
}

func (receiver *setGetFlusherMock) Get(key string) (interface{}, bool) {
	return getMock(key)
}

func (receiver *setGetFlusherMock) Set(key string, value interface{}, duration time.Duration) {
	setMock(key, value, duration)
}

func (receiver *setGetFlusherMock) Flush() {
	flushMock()
}

func (receiver *setGetFlusherMock) Items() map[string]cache.Item {
	return itemsMock()
}

func TestStorageService_Get(t *testing.T) {

	var verifyCalled bool
	var getKey string

	type fields struct {
		cache SetGetFlusher
	}
	type args struct {
		key     string
		getMock func(key string) (interface{}, bool)
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantCalled bool
		want       Item
	}{
		{
			name: "should return error when value does not exist",
			fields: fields{
				cache: &setGetFlusherMock{},
			},
			args: args{
				key: "valid",
				getMock: func(key string) (interface{}, bool) {
					getKey = key
					verifyCalled = true
					return nil, false
				},
			},
			wantErr:    true,
			wantCalled: true,
			want:       Item{},
		},
		{
			name: "should return value when getter returns value",
			fields: fields{
				cache: &setGetFlusherMock{},
			},
			args: args{
				key: "valid",
				getMock: func(key string) (interface{}, bool) {
					getKey = key
					verifyCalled = true
					return Item{
						BattlePoint: 123,
						MinePoint:   321,
					}, true
				},
			},
			wantErr:    false,
			wantCalled: true,
			want: Item{
				BattlePoint: 123,
				MinePoint:   321,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getMock = tt.args.getMock

			receiver := StorageService{
				cache: tt.fields.cache,
			}
			err2, value := receiver.Get(tt.args.key)
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("Retrieve() got = %v, want %v", value, tt.want)
			}
			if tt.wantCalled != verifyCalled {
				t.Errorf("call for Get() method on Getter does not match mock. wanted: %v actual: %v", tt.wantCalled, verifyCalled)
			}
			if tt.wantCalled && !reflect.DeepEqual(getKey, "valid") {
				t.Errorf("call for Get() method on Getter does not match mock key. wanted: %v actual: %v", "valid", getKey)
			}
			verifyCalled = false
			getKey = ""
		})
	}
}

func TestStorageService_Set(t *testing.T) {

	var verifyCalled bool
	var setKey string
	var setValue string
	var setDuration time.Duration

	type fields struct {
		cache SetGetFlusher
	}
	type args struct {
		key      string
		value    interface{}
		duration time.Duration
		setMock  func(key string, value interface{}, d time.Duration)
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCalled bool
	}{
		{
			name: "should call set mock",
			fields: fields{
				cache: &setGetFlusherMock{},
			},
			args: args{
				key:      "valid",
				value:    "request",
				duration: 1 * time.Hour,
				setMock: func(key string, value interface{}, d time.Duration) {
					setKey = key
					setValue = value.(string)
					verifyCalled = true
					setDuration = d
				},
			},
			wantCalled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setMock = tt.args.setMock

			receiver := StorageService{
				cache: tt.fields.cache,
			}
			receiver.Set(tt.args.key, tt.args.value)
			if tt.wantCalled != verifyCalled {
				t.Errorf("call for Set() method on Setter does not match mock. wanted: %v actual: %v", tt.wantCalled, verifyCalled)
			}
			if !tt.wantCalled {
				return
			}
			if !reflect.DeepEqual(setKey, tt.args.key) {
				t.Errorf("call for Set() method on Setter does not match mock key. wanted: %v actual: %v", tt.args.key, setKey)
			}
			if !reflect.DeepEqual(setValue, tt.args.value) {
				t.Errorf("call for Set() method on Setter does not match mock value. wanted: %v actual: %v", tt.args.value, setValue)
			}
			if !reflect.DeepEqual(setDuration.Milliseconds(), tt.args.duration.Milliseconds()) {
				t.Errorf("call for Set() method on Setter does not match mock duration. wanted: %v actual: %v", tt.args.duration, setDuration)
			}
			verifyCalled = false
			setKey = ""
			setValue = ""
			setDuration = 99
		})
	}
}

func TestStorageService_Flush(t *testing.T) {

	var verifyCalled bool

	type fields struct {
		cache SetGetFlusher
	}
	type args struct {
		flushMock func()
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCalled bool
	}{
		{
			name: "should call set mock",
			fields: fields{
				cache: &setGetFlusherMock{},
			},
			args: args{
				flushMock: func() {
					verifyCalled = true
				},
			},
			wantCalled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flushMock = tt.args.flushMock

			receiver := StorageService{
				cache: tt.fields.cache,
			}
			receiver.Flush()
			if tt.wantCalled != verifyCalled {
				t.Errorf("call for Flush() method on Flusher does not match mock. wanted: %v actual: %v", tt.wantCalled, verifyCalled)
			}
			verifyCalled = false
		})
	}
}