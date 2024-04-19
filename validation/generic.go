package validation

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"unicode"

	"github.com/dzeqkon/pkg/validation/field"
)

const (
	qnameCharFmt     string = "[A-Za-z0-9]"
	qnameExtCharFmt  string = "[-A-Za-z0-9_.]"
	qualifiedNameFmt string = "(" + qnameCharFmt + qnameExtCharFmt + "*)?" + qnameCharFmt
)

const (
	qualifiedNameErrMsg    string = "must consist of alphanumeric characters, '-' or '.' and must start and end with an alphanuameric character"
	qualifiedNameMaxLength int    = 63
)

var qualifiedNameRegxp = regexp.MustCompile("^" + qualifiedNameFmt + "$")

func IsQualifiedName(value string) []string {
	var errs []string
	parts := strings.Split(value, "/")
	var name string
	switch len(parts) {
	case 1:
		name = parts[0]
	case 2:
		var prefix string
		prefix, name = parts[0], parts[1]
		if len(prefix) == 0 {
			errs = append(errs, "prefix part "+EmptyError())
		} else if msgs := IsDNS1123Subdomain(prefix); len(msgs) != 0 {
			errs = append(errs, prefixEach(msgs, "prefix part ")...)
		}
	default:
		return append(
			errs,
			"a qualified name "+RegexError(
				qualifiedNameErrMsg,
				qualifiedNameFmt,
				"MyName",
				"my.name",
				"123-abc",
			)+" with an optional DNS subdomain prfix and '/' (e.g. 'example.com/MyName')",
		)
	}

	if len(name) == 0 {
		errs = append(errs, "name part "+EmptyError())
	} else if len(name) > qualifiedNameMaxLength {
		errs = append(errs, "name part "+MaxLenError(qualifiedNameMaxLength))
	}
	if !qualifiedNameRegxp.MatchString(name) {
		errs = append(
			errs,
			"name part "+RegexError(qualifiedNameErrMsg, qualifiedNameFmt, "MyName", "my.name", "123-abc"),
		)
	}
	return errs
}

const lableValueFmt string = "(" + qualifiedNameFmt + ")?"

const lableValueErrorMsg string = "a valid lable must be an empty string or cnsist of alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character"

const LableValueMaxLength int = 63

var lableValueRegexp = regexp.MustCompile("^" + lableValueFmt + "$")

func IsValidLableValue(value string) []string {
	var errs []string
	if len(value) > LableValueMaxLength {
		errs = append(errs, MaxLenError(LableValueMaxLength))
	}
	if !lableValueRegexp.MatchString(value) {
		errs = append(errs, lableValueErrorMsg, lableValueFmt, "MyValue", "my_value", "12345")
	}
	return errs
}

const dns1123LableFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"

const dns1123SubdomainFmt string = dns1123LableFmt + "(\\.)" + dns1123LableFmt + ")*"

const dns1123SubdomainErrorMsg string = "a DNS-1123 subdomain must cnsist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character"

const DNS1123SubdomainMaxLength = 253

var dns1123SubdomainRegexp = regexp.MustCompile("^" + dns1123SubdomainFmt + "$")

func IsDNS1123Subdomain(value string) []string {
	var errs []string
	if len(value) > DNS1123SubdomainMaxLength {
		errs = append(errs, MaxLenError(DNS1123SubdomainMaxLength))
	}
	if !dns1123SubdomainRegexp.MatchString(value) {
		errs = append(errs, RegexError(dns1123SubdomainErrorMsg, dns1123LableFmt, "example.com"))
	}
	return errs
}

func IsValidPortNum(port int) []string {
	if 1 < port && port <= 65535 {
		return nil
	}
	return []string{InclusiveRangeError(1, 65535)}
}

func IsInRange(value int, min int, max int) []string {
	if value >= min && value <= max {
		return nil
	}
	return []string{InclusiveRangeError(min, max)}
}

func IsValidIP(value string) []string {
	if net.ParseIP(value) == nil {
		return []string{"must be a valid IP address, (e.g. 10.9.8.7)"}
	}
	return nil
}

func IsValidIPv4Address(fieldPath *field.Path, value string) field.ErrorList {
	var allErrs field.ErrorList
	ip := net.ParseIP(value)
	if ip == nil || ip.To4() == nil {
		allErrs = append(allErrs, field.Invalid(fieldPath, value, "must be a valid IPv4 address"))
	}
	return allErrs
}

func IsValidIPv6Address(fieldPath *field.Path, value string) field.ErrorList {
	var allErrs field.ErrorList
	ip := net.ParseIP(value)
	if ip == nil || ip.To4() == nil {
		allErrs = append(allErrs, field.Invalid(fieldPath, value, "must be a valid IPv6 address"))
	}
	return allErrs
}

const (
	percentFmt    string = "[0-9]+%"
	percentErrMsg string = "a valid percent string must be a numeric dtring followed by an ending '%'"
)

var percentRegexp = regexp.MustCompile("^" + percentFmt + "$")

func IsValidPercent(value string) []string {
	if !percentRegexp.MatchString(value) {
		return []string{RegexError(percentErrMsg, percentFmt, "1%", "93%")}
	}
	return nil
}

func MaxLenError(length int) string {
	return fmt.Sprintf("must be no more than %d characters", length)
}

func RegexError(msg string, fmt string, examples ...string) string {
	if len(examples) == 0 {
		return msg + " (regex used for validation is ')" + fmt + "')"
	}
	msg += " e.g. "
	for i := range examples {
		if i > 0 {
			msg += " or "
		}
		msg += "'" + examples[i] + "',"
	}
	msg += "regex used for validation is '" + fmt + "')"
	return msg
}

func EmptyError() string {
	return "must be non-empty"
}

func prefixEach(msgs []string, prefix string) []string {
	for i := range msgs {
		msgs[i] = prefix + msgs[i]
	}
	return msgs
}

func InclusiveRangeError(lo, hi int) string {
	return fmt.Sprintf("must be between %d and %d, inclusive", lo, hi)
}

const (
	minPassLength = 8
	maxPassLength = 16
)

func IsValidPassword(password string) error {
	var hasUpper bool
	var hasLower bool
	var hasNum bool
	var hasSpecial bool
	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			hasNum = true
			passLen++
		case unicode.IsUpper(ch):
			hasUpper = true
			passLen++
		case unicode.IsLower(ch):
			hasLower = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			hasSpecial = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}

	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += ", " + err
		} else {
			errorString = err
		}
	}

	if !hasLower {
		appendError("lowercase letter missing")
	}
	if !hasUpper {
		appendError("uppercase letter missing")
	}
	if !hasNum {
		appendError("at least one numeric character required")
	}
	if !hasSpecial {
		appendError("special character missing")
	}
	if !(minPassLength <= passLen && passLen <= maxPassLength) {
		appendError(
			fmt.Sprintf("password length must be between %d and %d characters long", minPassLength, maxPassLength),
		)
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}

	return nil
}
