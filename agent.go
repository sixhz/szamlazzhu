package szamlazzhu

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Agent is an authenticated agent that can interact with the szamlazz.hu API endpoint
// https://docs.szamlazz.hu
type Agent struct {
	// One of username+password or token is required. Other must be zero value
	username string // Authentication username
	password string // Authentication password
	token    string // Authentication token

	http *http.Client // Used for storing state between calls
}

// NewUserAgent creates a new Agent instance, based on username authentication
func NewUserAgent(username string, password string) *Agent {
	var a Agent
	a.username = username
	a.password = password
	a.init()
	return &a
}

// NewAgentUser creates a new Agent instance, based on token authentication
func NewTokenAgent(token string) *Agent {
	var a Agent
	a.token = token
	a.init()
	return &a
}

// init performs the actual intialization of an empty Agent
func (a *Agent) init() {
	a.http = new(http.Client)
	a.http.Timeout = 60 * time.Second
	// SzamlaAgent documentation specifically requires clients to handle cookies.
	// Using "nil" as PublicSuffixList is safe as long this http client is only used by our API
	a.http.Jar, _ = cookiejar.New(nil)
}

// GenerateInvoice creates a new invoice (or other document)
// https://docs.szamlazz.hu/#generating-invoices
func (a *Agent) GenerateInvoice(req Xmlszamla) (Xmlszamlavalasz, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token
	req.ValaszVerzio = 2

	// execute rpc
	var rpl Xmlszamlavalasz
	err := a.rpcCall(&req, "action-xmlagentxmlfile", &rpl)
	if err != nil {
		return Xmlszamlavalasz{}, err
	}

	// check for xml-level error
	if !rpl.Sikeres {
		return Xmlszamlavalasz{}, &SzamlazzhuError{rpl.Hibakod, rpl.Hibauzenet}
	}

	return rpl, nil
}

var stornoRegex = regexp.MustCompile(`^\s*xmlagentresponse=DONE;(.+?)\s*$`)

// ReverseInvoice reverses an existing invoice and returns the reversing document number
// https://docs.szamlazz.hu/#reversing-an-invoice-storno
func (a *Agent) ReverseInvoice(req Xmlszamlast) (string, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl stringContainer
	err := a.rpcCall(&req, "action-szamla_agent_st", &rpl)
	if err != nil {
		return "", err
	}

	// parse string reply
	rplMatch := stornoRegex.FindStringSubmatch(rpl.s)
	if rplMatch == nil {
		return "", fmt.Errorf("server replied garbage: %q", rpl.s)
	}

	return rplMatch[1], nil
}

// RegisterCredit updates payment entries for an invoice
// https://docs.szamlazz.hu/#registering-a-credit-entry
func (a *Agent) RegisterCredit(req Xmlszamlakifiz) error {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl stringContainer
	err := a.rpcCall(&req, "action-szamla_agent_kifiz", &rpl)
	if err != nil {
		return err
	}

	// check string reply
	if strings.TrimSpace(rpl.s) != "xmlagentresponse=DONE" {
		return fmt.Errorf("server replied garbage: %q", rpl.s)
	}

	return nil
}

// QueryInvoicePdf downloads the PDF representation of an invoice
// https://docs.szamlazz.hu/#querying-the-invoice-pdf
func (a *Agent) QueryInvoicePdf(req Xmlszamlapdf) (Xmlszamlavalasz, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token
	req.ValaszVerzio = 2

	// execute rpc
	var rpl Xmlszamlavalasz
	err := a.rpcCall(&req, "action-szamla_agent_pdf", &rpl)
	if err != nil {
		return Xmlszamlavalasz{}, err
	}

	// check for xml-level error
	if !rpl.Sikeres {
		return Xmlszamlavalasz{}, &SzamlazzhuError{rpl.Hibakod, rpl.Hibauzenet}
	}

	return rpl, nil
}

// QueryInvoiceXml downloads invoice details
// https://docs.szamlazz.hu/#querying-the-invoice-pdf
func (a *Agent) QueryInvoiceXml(req Xmlszamlaxml) (Szamla, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl Szamla
	err := a.rpcCall(&req, "action-szamla_agent_xml", &rpl)
	if err != nil {
		return Szamla{}, err
	}

	return rpl, nil
}

// DeleteProforma removes a proforma invoice (díjbekérő) from szamlazz.hu
// https://docs.szamlazz.hu/#deleting-a-pro-forma-invoice
func (a *Agent) DeleteProforma(req Xmlszamladbkdel) (Xmlszamladbkdelvalasz, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl Xmlszamladbkdelvalasz
	err := a.rpcCall(&req, "action-szamla_agent_dijbekero_torlese", &rpl)
	if err != nil {
		return Xmlszamladbkdelvalasz{}, err
	}

	// check for xml-level error
	if !rpl.Sikeres {
		return Xmlszamladbkdelvalasz{}, &SzamlazzhuError{rpl.Hibakod, rpl.Hibauzenet}
	}

	return rpl, nil
}

// GenerateReceipt creates a new receipt
// https://docs.szamlazz.hu/#generating-a-receipt
func (a *Agent) GenerateReceipt(req Xmlnyugtacreate) (Xmlnyugtavalasz, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl Xmlnyugtavalasz
	err := a.rpcCall(&req, "action-szamla_agent_nyugta_create", &rpl)
	if err != nil {
		return Xmlnyugtavalasz{}, err
	}

	// check for xml-level error
	if !rpl.Sikeres {
		return Xmlnyugtavalasz{}, &SzamlazzhuError{rpl.Hibakod, rpl.Hibauzenet}
	}

	return rpl, nil
}

// StornoReceipt is an alias of ReverseReceipt for compatibility reasons.
// Deprecated!
func (a *Agent) StornoReceipt(req Xmlnyugtast) (Xmlnyugtavalasz, error) {
	return a.ReverseReceipt(req)
}

// ReverseReceipt reverses an existing receipt
// https://docs.szamlazz.hu/#reversing-a-receipt-storno
func (a *Agent) ReverseReceipt(req Xmlnyugtast) (Xmlnyugtavalasz, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl Xmlnyugtavalasz
	err := a.rpcCall(&req, "action-szamla_agent_nyugta_storno", &rpl)
	if err != nil {
		return Xmlnyugtavalasz{}, err
	}

	// check for xml-level error
	if !rpl.Sikeres {
		return Xmlnyugtavalasz{}, &SzamlazzhuError{rpl.Hibakod, rpl.Hibauzenet}
	}

	return rpl, nil
}

// QueryReceipt downloads an existing receipt
// https://docs.szamlazz.hu/#querying-a-receipt
func (a *Agent) QueryReceipt(req Xmlnyugtaget) (Xmlnyugtavalasz, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl Xmlnyugtavalasz
	err := a.rpcCall(&req, "action-szamla_agent_nyugta_get", &rpl)
	if err != nil {
		return Xmlnyugtavalasz{}, err
	}

	// check for xml-level error
	if !rpl.Sikeres {
		return Xmlnyugtavalasz{}, &SzamlazzhuError{rpl.Hibakod, rpl.Hibauzenet}
	}

	return rpl, nil
}

// SendReceipt sends an existing receipt via e-mail
// https://docs.szamlazz.hu/#sending-a-receipt
func (a *Agent) SendReceipt(req Xmlnyugtasend) (Xmlnyugtasendvalasz, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl Xmlnyugtasendvalasz
	err := a.rpcCall(&req, "action-szamla_agent_nyugta_send", &rpl)
	if err != nil {
		return Xmlnyugtasendvalasz{}, err
	}

	// check for xml-level error
	if !rpl.Sikeres {
		return Xmlnyugtasendvalasz{}, &SzamlazzhuError{rpl.Hibakod, rpl.Hibauzenet}
	}

	return rpl, nil
}

// QueryTaxpayer requests taxpayer information from NAV via szamlazz.hu
// https://docs.szamlazz.hu/#querying-taxpayers
func (a *Agent) QueryTaxpayer(req Xmltaxpayer) (QueryTaxpayerResponse, error) {

	// hardcoded fields
	req.Felhasznalo = a.username
	req.Jelszo = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl QueryTaxpayerResponse
	err := a.rpcCall(&req, "action-szamla_agent_taxpayer", &rpl)
	if err != nil {
		return QueryTaxpayerResponse{}, err
	}

	// check for xml-level error
	if rpl.ResultFuncCode != "OK" {
		return QueryTaxpayerResponse{}, fmt.Errorf("%s %s %s", rpl.ResultFuncCode, rpl.ResultErrorCode, rpl.ResultMessage)
	}

	return rpl, nil
}

// SupplierStatus describes the state of a supplier account
type SupplierStatus int

const (
	Unknown          SupplierStatus = iota
	NewCompany                      // New supplier account creation has started
	NewCompanyResend                // New supplier account creation was already started, e-mail resent
	NewConnect                      // Supplier has existing account, access request sent
	NewConnectResend                // Supplier has existing account, access was already requested, e-mail resent
)

// agentresponseRegex parses a text-based reply for SupplierAccount calls
var agentresponseRegex = regexp.MustCompile(`^\s*xmlagentresponse=(.+?)\s*$`)

// SupplierAccount creates or requests access to a supplier's szamlazz.hu account
// https://docs.szamlazz.hu/#self-billing
func (a *Agent) SupplierAccount(req Xmlcegmb) (SupplierStatus, error) {
	// Possible text responses found by us:
	// xmlagentresponse=DONE
	// xmlagentresponse=Már létező fiók, nincs fiókgazdája. Fiókgazdai meghívó (megbízott számlakibocsátás) email újraküldve.
	// xmlagentresponse=Már létező fiók fiókgazdával. Csatlakozási kérelmet (ASK) küldtünk a fiókgazdának.
	// xmlagentresponse=Már létező fiók fiókgazdával. A csatlakozási kérelem emailt (ASK) újraküldtük a fiókgazdának.

	// hardcoded fields
	req.LoginName = a.username
	req.Password = a.password
	req.SzamlaAgentKulcs = a.token

	// execute rpc
	var rpl stringContainer
	err := a.rpcCall(&req, "action-agent_ceg_mb", &rpl)
	if err != nil {
		return Unknown, err
	}

	// parse string reply
	rplMatch := agentresponseRegex.FindStringSubmatch(rpl.s)
	if rplMatch == nil {
		return Unknown, fmt.Errorf("server replied garbage: %q", rpl.s)
	}

	switch rplMatch[1] {
	case "DONE":
		return NewCompany, nil
	case "Már létező fiók, nincs fiókgazdája. Fiókgazdai meghívó (megbízott számlakibocsátás) email újraküldve.":
		return NewCompanyResend, nil
	case "Már létező fiók fiókgazdával. Csatlakozási kérelmet (ASK) küldtünk a fiókgazdának.":
		return NewConnect, nil
	case "Már létező fiók fiókgazdával. A csatlakozási kérelem emailt (ASK) újraküldtük a fiókgazdának.":
		return NewConnectResend, nil
	default:
		return Unknown, fmt.Errorf("unknown state xmlagentresponse=%s", rplMatch[1])
	}
}

// apiEndpoint is the szamlazz.hu API URL
const apiEndpoint = "https://www.szamlazz.hu/szamla/"

// postAsFile will uploads a file as HTTP web form file upload
func (a *Agent) postAsFile(field string, filename string, data []byte) ([]byte, error) {
	// Package data into a multipart message and buffer it
	var buffer bytes.Buffer
	msg := multipart.NewWriter(&buffer)
	i, _ := msg.CreateFormFile(field, filename)
	i.Write(data)
	msg.Close()

	// Send the HTTP request to the Szamlazz.hu server
	response, err := a.http.Post(apiEndpoint, msg.FormDataContentType(), &buffer)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Szamlazz.hu API always replies 200 OK. Anything else means the request was misrouted or lost.
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("received HTTP %d %s during file upload", response.StatusCode, response.Status)
	}

	// Szamlazz.hu may indicate errors via non-standard HTTP custom headers
	if ecode := response.Header.Get("szlahu_error_code"); ecode != "" {
		c, _ := strconv.Atoi(ecode)
		s, _ := url.QueryUnescape(response.Header.Get("szlahu_error"))
		return nil, &SzamlazzhuError{c, s}
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP response %w", err)
	}

	return body, nil
}

// errRegex matches stack-trace embedded textual error messages from szamlazz.hu
var errRegex = regexp.MustCompile(`(?s)^\s*\[ERR\](.*?)----------`)

// rpcCall performs an RPC request to szamlazz.hu API
// This function takes care of the encoding/decoding of message structs, as well as submitting the encoded data
// Note that the request should, the reply _must_ be a pointer for correct behavior
func (a *Agent) rpcCall(request Message, field string, reply Message) error {
	// Convert incoming data to XML representation
	datareq, err := xml.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal RPC message: %w", err)
	}

	// Communicate with szamlazz.hu
	dataresp, err := a.postAsFile(field, "agent.xml", datareq)
	if err != nil {
		return err
	}

	// szamlazz.hu may throw a stacktrace for any request, we can find these errors via regex
	if szlaerr := errRegex.FindSubmatch(dataresp); szlaerr != nil {
		return &SzamlazzhuError{-1, string(szlaerr[1])}
	}

	// some endpoints return string instead of XML, use stringContainer to encapsulate them
	if str, ok := reply.(*stringContainer); ok {
		str.s = string(dataresp)
		return nil
	}

	// attempt to understand the xml reply
	err = xml.Unmarshal(dataresp, reply)
	if err != nil {
		// some endpoints return xmlszamlavalasz upon error instead of the contracted XSD schema
		var xmlerror Xmlszamlavalasz
		if errDecode := xml.Unmarshal(dataresp, &xmlerror); errDecode == nil && !xmlerror.Sikeres {
			return &SzamlazzhuError{xmlerror.Hibakod, xmlerror.Hibauzenet}
		}

		return fmt.Errorf("failed to unmarshal RPC reply: %w", err)
	}

	return nil
}
