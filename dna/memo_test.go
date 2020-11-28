package dna

import (
	"testing"

	"github.com/blocktree/bitshares-adapter/types"
)

func TestDecrypt(t *testing.T) {
	type args struct {
		msg     string
		fromPub string
		toPub   string
		wif     string
		nonce   uint64
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// {
		// 	name: "test1",
		// 	want: "hello boy",
		// 	args: args{
		// 		msg:     "17d0ac3874548d7c4ef56236698d719e",
		// 		nonce:   5577006791947779410,
		// 		fromPub: "DNA6fpcoYK72BxsYYRwcBEPGuVoGhy2Yki2YCNfnCZCYxL5xp56Hh",
		// 		toPub:   "DNA6icdz8dWibXRz8PcDn9RMupFkPbwHQ4toHxP8UmLm2hDtMHUKr",
		// 		wif:     "",
		// 	},
		// },
		// {
		// 	name: "test2",
		// 	want: "hellowrd111122223333",
		// 	args: args{
		// 		msg:     "5a9436f43ba1299afca79ebbf9698b08ace004eba9c17a789d158b74d195ed84",
		// 		nonce:   5405926024387204188,
		// 		fromPub: "DNA8Xi4HwBaUAkYD5wT7jRaAS1B2NUqw8rzE29Vdaww5zvtsse7tW",
		// 		toPub:   "DNA6fpcoYK72BxsYYRwcBEPGuVoGhy2Yki2YCNfnCZCYxL5xp56Hh",
		// 		wif:     "",
		// 	},
		// },
		{
			name: "test3",
			want: "123456789012",
			args: args{
				msg:     "d08f99ecaab4bb6eab51bd39d3c65462",
				nonce:   13260572831089785859,
				fromPub: "DNA6fpcoYK72BxsYYRwcBEPGuVoGhy2Yki2YCNfnCZCYxL5xp56Hh",
				toPub:   "DNA6fpcoYK72BxsYYRwcBEPGuVoGhy2Yki2YCNfnCZCYxL5xp56Hh",
				wif:     "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf types.Buffer
			buf.FromString(tt.args.msg)
			got, err := Decrypt(buf, tt.args.fromPub, tt.args.toPub, tt.args.nonce, tt.args.wif)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
