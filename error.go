package szamlazzhu

import "fmt"

type SzamlazzhuError struct {
	c int
	e string
}

func (e *SzamlazzhuError) Error() string {
	return fmt.Sprintf("szamlazz.hu error %d: %s", e.c, e.e)
}

// Configuration errors - you may need to change
const (
	ErrorSystemMaintenance    = 1   // System Maintenance, please try again in a few minutes.
	ErrorFailedLogin          = 3   // Login error. Invalid login name, password or token.
	ErrorEinvoiceUnauthorized = 54  // The issuing of e-invoices is not permitted for your account.
	ErrorAccountProblem       = 136 // For some reason you cannot use Számla Agent. Please log in to the Számlázz.hu system via a browser.
	ErrorMultiUser            = 164 // The username-password combination is ambigous due to multi-company access. Please use Token authentication.
	ErrorInvalidPrefix        = 202 // The provided invoice prefix cannot be used. You need to login to the website and add the desired prefixes before you can use them with Számla Agent.
)

// Structural errors - these may be bugs in the binding package, consider reporting them.
const (
	ErrorXMLFile         = 53  // Missing XML file. RPC structure error
	ErrorXMLRead         = 57  // XML reading error. There is an error in the sent XML file.
	ErrorSessionConflict = 135 // Conflict in session data
)

// Data errors - you provided information that does not pass szamlazz.hu validations.
const (
	ErrorInvalidNettoErtek1  = 259 // The NET value of the item is not correct
	ErrorInvalidAfa1         = 260 // The VAT value of the item is not correct
	ErrorInvalidBruttoErtek1 = 261 // The GROSS value of the item is not correct
	ErrorInvalidNettoErtek2  = 262 // The NET value of the item is not correct
	ErrorInvalidAfa2         = 263 // The VAT value of the item is not correct
	ErrorInvalidBruttoErtek2 = 264 // The GROSS value of the item is not correct
)
