package mail

import (
	"net/mail"
	"strings"

	"github.com/ProtonMail/go-rfc5322"
	"github.com/linanh/go-message/charset"
)

// Address represents a single mail address.
// The type alias ensures that a net/mail.Address can be used wherever an
// Address is expected
type Address = mail.Address

func init() {
	rfc5322.CharsetReader = charset.Reader
}

func formatAddressList(l []*Address) string {
	formatted := make([]string, len(l))
	for i, a := range l {
		formatted[i] = a.String()
	}
	return strings.Join(formatted, ", ")
}

// ParseAddress parses a single RFC 5322 address, e.g. "Barry Gibbs <bg@example.com>"
// Use this function only if you parse from a string, if you have a Header use
// Header.AddressList instead
func ParseAddress(address string) (*Address, error) {
	addrList, err := rfc5322.ParseAddressList(address)
	var addr *Address
	if len(addrList) > 0 {
		addr = addrList[0]
	}
	return addr, err
}

// ParseAddressList parses the given string as a list of addresses.
// Use this function only if you parse from a string, if you have a Header use
// Header.AddressList instead
func ParseAddressList(list string) ([]*Address, error) {
	return rfc5322.ParseAddressList(list)
}
