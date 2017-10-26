package helpers

import (
	"flag"
	"log"
	"os"
	"text/tabwriter"
	"github.com/dim13/unifi"
	"strings"
)

var (
	num    = flag.String("num", "1", "The number of the new vouchers")
	multi  = flag.String("multi", "1", "If 0 is the multi-voucher")
	minute = flag.String("minute", "120", "Duration of the voucher in minutes")
	note   = flag.String("note", "Auto-generated voucher", "Note of the voucher")

	host = flag.String("host", "10.11.0.177", "Controller hostname")
	user = flag.String("user", "Admin", "Controller username")
	pass = flag.String("pass", "ejxnt5*", "Controller password")
)


func GenerateVoucher() string {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 3, ' ', 0)
	defer w.Flush()

	flag.Parse()
	version := new(int)
	siteid := new(string)
	*version = 4
	*siteid = "default"

	u, err := unifi.Login(*user, *pass, *host, *siteid, *version)
	if err != nil {
		log.Fatal("Login returned error: ", err)
	}
	defer u.Logout()

	jsonData := unifi.NewVoucher{
		Cmd:          "create-voucher",
		Expire:       "custom",
		ExpireNumber: *minute,
		ExpireUnit:   "1",
		N:            *num,
		Note:         *note,
		Quota:        *multi,
	}

	res, err := u.NewVoucher(jsonData)
	if err != nil {
		log.Fatalln(err)
	}

	ct := res[0].CreateTime

	vouchers, err := u.VoucherMap()

	if err != nil {
		log.Fatalln(err)
	}

	p := []string {}

	for _, v := range vouchers {
		if ct == v.CreateTime {
			p = append(p, v.Code)
		}
	}

	return strings.Join(p, ",")
}