package resource

import "github.com/aelsabbahy/goss/system"

type DNS struct {
	Host        string   `json:"-"`
	Resolveable bool     `json:"resolveable"`
	Addrs       []string `json:"addrs,omitempty"`
	Timeout     int64    `json:"timeout"`
}

func (d *DNS) ID() string      { return d.Host }
func (d *DNS) SetID(id string) { d.Host = id }

func (d *DNS) Validate(sys *system.System) []TestResult {
	sysDNS := sys.NewDNS(d.Host, sys)
	sysDNS.SetTimeout(d.Timeout)

	var results []TestResult

	results = append(results, ValidateValue(d, "resolveable", d.Resolveable, sysDNS.Resolveable))
	if !d.Resolveable {
		return results
	}
	if len(d.Addrs) > 0 {
		results = append(results, ValidateValues(d, "addrs", d.Addrs, sysDNS.Addrs))
	}

	return results
}

func NewDNS(sysDNS system.DNS) *DNS {
	host := sysDNS.Host()
	addrs, _ := sysDNS.Addrs()
	resolveable, _ := sysDNS.Resolveable()
	return &DNS{
		Host:        host,
		Addrs:       addrs,
		Resolveable: resolveable.(bool),
		Timeout:     500,
	}
}
