package usecase

import (
	"fmt"
	"time"

	"github.com/rymiyamoto/affiliate-api/flag"
	"github.com/rymiyamoto/affiliate-api/service"
)

type (
	IRakuten interface {
		Exec(f *flag.Affiliate) error
	}

	Rakuten struct {
		rakutenClientService service.IRakutenClient
	}
)

// NewRakuten initialize
func NewRakuten() IRakuten {
	return &Rakuten{
		rakutenClientService: service.NewRakutenClient(),
	}
}

// Exec 実行
func (u *Rakuten) Exec(f *flag.Affiliate) error {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if err := u.rakutenClientService.GetProduct(); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}
