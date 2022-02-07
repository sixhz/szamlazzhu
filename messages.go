package szamlazzhu

// This file contains the structs describing the xml messages sent and received from szamlazz.hu endpoints
// Note: We want to avoid Hungarian names in our source, but we have to make an exception here.
// The szamlazz.hu API use hard-coded hungarian tags in the XML format, which we will blindly follow.

// Empty is a dummy type that we use for passing in tags data reflections. Ideal, because it does not allocate memory
type Empty struct{}

// Message is an alias of empty interface which we use for passing the endpoint descriptor structs
type Message interface{}

// stringContainer is a magic struct that is used to return simple strings in place of XML structures
type stringContainer struct {
	s string
}

// Xmlszamla describes the request of creating new invoices
type Xmlszamla struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlszamla xmlszamla"`

	// Authentication (Username+Password or AgentKey required)

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`

	ESzamla          bool   `xml:"beallitasok>eszamla"`                    // Whether to issue electronically signed invoices
	KulcstartoJelszo string `xml:"beallitasok>kulcstartojelszo,omitempty"` // Password for electronic signature keyring
	SzamlaLetoltes   bool   `xml:"beallitasok>szamlaLetoltes"`             // Whether to return the created invoice image
	ValaszVerzio     uint   `xml:"beallitasok>valaszVerzio"`               // Format of reply. 1: mixed 2: xml-encapsulated

	KeltDatum              Date    `xml:"fejlec>keltDatum,omitempty"`              // Date of creation (YYYY-MM-DD)
	TeljesitesDatum        Date    `xml:"fejlec>teljesitesDatum"`                  // Date of fullfillment (YYYY-MM-DD)
	FizetesiHataridoDatum  Date    `xml:"fejlec>fizetesiHataridoDatum"`            // Payment deadline (YYYY-MM-DD)
	FizMod                 string  `xml:"fejlec>fizmod"`                           // Payment type
	Penznem                string  `xml:"fejlec>penznem"`                          // Currency
	SzamlaNyelve           string  `xml:"fejlec>szamlaNyelve"`                     // Language (hu/de/en/it/fr/ro/sk/hr)
	Megjegyzes             string  `xml:"fejlec>megjegyzes"`                       // Comment
	ArfolyamBank           string  `xml:"fejlec>arfolyamBank,omitempty"`           // Bank of exchange (MNB)
	Arfolyam               float64 `xml:"fejlec>arfolyam,omitempty"`               // Exchange rate (0.0)
	RendelesSzam           string  `xml:"fejlec>rendelesSzam"`                     // Order number
	DijbekeroSzamlaszam    string  `xml:"fejlec>dijbekeroSzamlaszam,omitempty"`    // ID of Proforma invoice
	ElolegSzamla           bool    `xml:"fejlec>elolegszamla,omitempty"`           // Is advance invoice?
	VegSzamla              bool    `xml:"fejlec>vegszamla,omitempty"`              // Is settlement invoice?
	HelyesbitoSzamla       bool    `xml:"fejlec>helyesbitoszamla,omitempty"`       // Is correction invoice?
	HelyesbitettSzamlaszam bool    `xml:"fejlec>helyesbitettSzamlaszam,omitempty"` // No of invoice corrected
	Dijbekero              bool    `xml:"fejlec>dijbekero,omitempty"`              // Is proforma invoice?
	Szallitolevel          bool    `xml:"fejlec>szallitolevel,omitempty"`          // Is delivery bill?
	SzamlaszamElotag       string  `xml:"fejlec>szamlaszamElotag,omitempty"`       // Invoice ID prefix
	Fizetve                bool    `xml:"fejlec>fizetve"`                          // Invoice has been paid?
	SzamlaSablon           string  `xml:"fejlec>szamlaSablon,omitempty"`           // Invoice format (SzlaMost/SzlaAlap/SzlaNoEnv/Szla8cm/SzlaTomb)
	Elonezetpdf            bool    `xml:"fejlec>elonezetpdf,omitempty"`            // Invoice format (SzlaMost/SzlaAlap/SzlaNoEnv/Szla8cm/SzlaTomb)

	Bank           string `xml:"elado>bank,omitempty"`           // Merchant Bank Name
	BankszamlaSzam string `xml:"elado>bankszamlaszam,omitempty"` // Merchant Bank Account
	EmailReplyTo   string `xml:"elado>emailReplyto,omitempty"`   // Reply-to header for notification e-mails
	EmailTargy     string `xml:"elado>emailTargy,omitempty"`     // Subject header for notification e-mails
	EmailSzoveg    string `xml:"elado>emailSzoveg,omitempty"`    // Content text of notification e-mails

	VevoNev                string `xml:"vevo>nev"`                          // Customer name
	VevoOrszag             string `xml:"vevo>orszag,omitempty"`             // Customer country
	VevoIrsz               string `xml:"vevo>irsz"`                         // Customer address zip code
	VevoTelepules          string `xml:"vevo>telepules"`                    // Customer address city
	VevoCim                string `xml:"vevo>cim"`                          // Customer address line
	VevoEmail              string `xml:"vevo>email,omitempty"`              // Customer e-mail address
	VevoSendEmail          bool   `xml:"vevo>sendEmail,omitempty"`          // Send notification e-mail to customer?
	VevoAdoszam            string `xml:"vevo>adoszam,omitempty"`            // Customer VAT number
	VevoAdoszamEU          string `xml:"vevo>adoszamEU,omitempty"`          // Customer VAT number
	VevoPostazasiNev       string `xml:"vevo>postazasiNev,omitempty"`       // Customer mail name
	VevoPostazasiOrszag    string `xml:"vevo>postazasiOrszag,omitempty"`    // Customer mail country
	VevoPostazasiIrsz      string `xml:"vevo>postazasiIrsz,omitempty"`      // Customer mail address zip code
	VevoPostazasiTelepules string `xml:"vevo>postazasiTelepules,omitempty"` // Customer mail address city
	VevoPostazasiCim       string `xml:"vevo>postazasiCim,omitempty"`       // Customer mail address line
	VevoAzonosito          string `xml:"vevo>azonosito,omitempty"`          // Customer unique ID
	VevoTelefonszam        string `xml:"vevo>telefonszam,omitempty"`        // Customer phone number
	VevoMegjegyzes         string `xml:"vevo>megjegyzes,omitempty"`         // Customer comment

	Tetelek []XmlszamlaTetel `xml:"tetelek>tetel"` // Invoice content
}

// XmlszamlaTetel describes a line item within an xmlszamla
type XmlszamlaTetel struct {
	Megnevezes       string  `xml:"megnevezes"`           // Item name
	Azonosito        string  `xml:"azonosito,omitempty"`  // Item identifier
	Mennyiseg        float64 `xml:"mennyiseg"`            // Item quantity
	MennyisegiEgyseg string  `xml:"mennyisegiEgyseg"`     // Quantity specifier
	NettoEgysegar    float64 `xml:"nettoEgysegar"`        // Net unit value
	AfaKulcs         string  `xml:"afakulcs"`             // VAT base (TAM/AAM/EU/EUK/MAA/F.AFA/K.AFA/0/5/7/18/19/20/25/27)
	NettoErtek       float64 `xml:"nettoErtek"`           // Net value
	AfaErtek         float64 `xml:"afaErtek"`             // VAT value
	BruttoErtek      float64 `xml:"bruttoErtek"`          // Gross value
	Megjegyzes       string  `xml:"megjegyzes,omitempty"` // Comment

	FokonyvGazdasagiEsem         string `xml:"tetelFokonyv>gazdasagiEsem,omitempty"`
	FokonyvGazdasagiEsemAfa      string `xml:"tetelFokonyv>gazdasagiEsemAfa,omitempty"`
	FokonyvArbevetelFokonyviSzam string `xml:"tetelFokonyv>arbevetelFokonyviSzam,omitempty"`
	FokonyvAfaFokonyviSzam       string `xml:"tetelFokonyv>afaFokonyviSzam,omitempty"`
	FokonyvElszDatumTol          Date   `xml:"tetelFokonyv>elszDatumTol,omitempty"`
	FokonyvElszDatumIg           Date   `xml:"tetelFokonyv>elszDatumIg,omitempty"`
}

// Xmlszamlavalasz describes the response of creating new invoices
type Xmlszamlavalasz struct {
	XMLName      Empty   `xml:"http://www.szamlazz.hu/xmlszamlavalasz xmlszamlavalasz"`
	Sikeres      bool    `xml:"sikeres"`
	Hibakod      int     `xml:"hibakod,omitempty"` // Note: this is string in the original XSD, parse it as int anyway
	Hibauzenet   string  `xml:"hibauzenet,omitempty"`
	Szamlaszam   string  `xml:"szamlaszam,omitempty"`
	SzamlaNetto  float64 `xml:"szamlanetto,omitempty"`
	SzamlaBrutto float64 `xml:"szamlabrutto,omitempty"`
	Kintlevoseg  float64 `xml:"kintlevoseg,omitempty"`
	Vevoifiokurl string  `xml:"vevoifiokurl,omitempty"`
	Pdf          []byte  `xml:"pdf,omitempty"`
}

// Xmlszamlast describes the request of reversing invoices
type Xmlszamlast struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlszamlast xmlszamlast"`

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`

	ESzamla          bool   `xml:"beallitasok>eszamla"`                    // Whether to issue electronically signed invoices
	KulcstartoJelszo string `xml:"beallitasok>kulcstartojelszo,omitempty"` // Password for electronic signature keyring

	SzamlaLetoltes bool `xml:"beallitasok>szamlaLetoltes"` // Whether to return the created invoice image
	ValaszVerzio   uint `xml:"beallitasok>valaszVerzio"`   // Format of reply. 1: mixed 2: xml-encapsulated

	Szamlaszam      string `xml:"fejlec>szamlaszam"`                // To be reversed invoice number
	KeltDatum       Date   `xml:"fejlec>keltDatum,omitempty"`       // Date of creation (YYYY-MM-DD)
	TeljesitesDatum Date   `xml:"fejlec>teljesitesDatum,omitempty"` // Date of fullfillment (YYYY-MM-DD)
	Tipus           string `xml:"fejlec>tipus,omitempty"`           // Type of reversal. Must be SS
	SzamlaSablon    string `xml:"fejlec>szamlaSablon,omitempty"`    // Invoice format (SzlaMost/SzlaAlap/SzlaNoEnv/Szla8cm/SzlaTomb)

	EmailReplyTo string `xml:"elado>emailReplyto,omitempty"` // Reply-to header for notification e-mails
	EmailTargy   string `xml:"elado>emailTargy,omitempty"`   // Subject header for notification e-mails
	EmailSzoveg  string `xml:"elado>emailSzoveg,omitempty"`  // Content text of notification e-mails

	Email string `xml:"vevo>email,omitempty"` // Customer e-mail address
}

// Xmlszamlakifiz describes the request of updating payment statuses
type Xmlszamlakifiz struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlszamlakifiz xmlszamlakifiz"`

	// Authentication (Username+Password or AgentKey required)
	Felhasznalo      string `xml:"beallitasok>felhasznalo"`      // API user
	Jelszo           string `xml:"beallitasok>jelszo"`           // API password
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs"` // API token

	Szamlaszam string `xml:"beallitasok>szamlaszam"` // Number of invoice to be updated
	Additiv    bool   `xml:"beallitasok>additiv"`    // Append payment records (true) or overwrite them (false)

	Kifizetes []XmlszamlakifizKifizetes `xml:"kifizetes"`
}

// XmlszamlakifizKifizetes describes a line item within an xmlszalakifiz
type XmlszamlakifizKifizetes struct {
	Datum  Date    `xml:"datum"`  // Date of creation (YYYY-MM-DD)
	Jogcim string  `xml:"jogcim"` // Method Of Payment (cash, credit card, etc)
	Osszeg float64 `xml:"osszeg"` // Net value
	Leiras string  `xml:"leiras"` // Description
}

// Xmlszamlapdf describes the request for downloading PDF invoices
type Xmlszamlapdf struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlszamlapdf xmlszamlapdf"`

	// Authentication (Username+Password or AgentKey required)
	Felhasznalo      string `xml:"felhasznalo,omitempty"`
	Jelszo           string `xml:"jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"szamlaagentkulcs,omitempty"`

	Szamlaszam   string `xml:"szamlaszam,omitempty"`   // Id of document to be downloaded
	RendelesSzam string `xml:"rendelesSzam,omitempty"` // Order number

	ValaszVerzio uint `xml:"valaszVerzio"` // Format of reply. 1: mixed 2: xml-encapsulated
}

// Xmlszamlaxml describes the request of downloading invoice data
type Xmlszamlaxml struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlszamlaxml xmlszamlaxml"`

	// Authentication (Username+Password or AgentKey required)
	Felhasznalo      string `xml:"felhasznalo,omitempty"`
	Jelszo           string `xml:"jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"szamlaagentkulcs,omitempty"`

	Szamlaszam   string `xml:"szamlaszam"`   // Id of document to be downloaded
	RendelesSzam string `xml:"rendelesSzam"` // Order number
	Pdf          bool   `xml:"pdf"`          // Whether to return the created invoice image
}

// Szamla describes the response of downloading invoice data
type Szamla struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/szamla szamla"`

	SzallitoId                int    `xml:"szallito>id"`
	SzallitoNev               string `xml:"szallito>nev"`
	SzallitoCimOrszag         string `xml:"szallito>cim>orszag"`
	SzallitoCimIrsz           string `xml:"szallito>cim>irsz"`
	SzallitoCimTelepules      string `xml:"szallito>cim>telepules"`
	SzallitoCimCim            string `xml:"szallito>cim>cim"`
	SzallitoPostacimOrszag    string `xml:"szallito>postacim>orszag,omitempty"`
	SzallitoPostacimIrsz      string `xml:"szallito>postacim>irsz,omitempty"`
	SzallitoPostacimTelepules string `xml:"szallito>postacim>telepules,omitempty"`
	SzallitoPostacimCim       string `xml:"szallito>postacim>cim,omitempty"`
	SzallitoAdoszam           string `xml:"szallito>adoszam"`
	SzallitoAdoszameu         string `xml:"szallito>adoszameu,omitempty"`
	SzallitoBankNev           string `xml:"szallito>bank>nev,omitempty"`
	SzallitoBankBankszamla    string `xml:"szallito>bank>bankszamla,omitempty"`

	AlapId            int     `xml:"alap>id"`
	AlapSzamlaszam    string  `xml:"alap>szamlaszam"`
	AlapForras        int     `xml:"alap>forras,omitempty"`
	AlapIktatoszam    string  `xml:"alap>iktatoszam"`
	AlapTipus         string  `xml:"alap>tipus"`
	AlapEszamla       int     `xml:"alap>eszamla"`
	AlapHivszamlaszam string  `xml:"alap>hivszamlaszam,omitempty"`
	AlapHivdijbekszam string  `xml:"alap>hivdijbekszam,omitempty"`
	AlapKelt          Date    `xml:"alap>kelt"`
	AlapTelj          Date    `xml:"alap>telj"`
	AlapFizh          Date    `xml:"alap>fizh"`
	AlapFizmod        string  `xml:"alap>fizmod"`
	AlapFizmodunified string  `xml:"alap>fizmodunified"` // ENUM: átutalás, készpénz, bankkártya, csekk, utánvét, ajándékutalvány, barion, barter, csoportos beszedés, OTP Simple, kompenzáció, kupon, PayPal, PayU, SZÉP kártya, utalvány, egyéb
	AlapKeszpenz      bool    `xml:"alap>keszpenz"`
	AlapRendelesszam  string  `xml:"alap>rendelesszam,omitempty"`
	AlapNyelv         string  `xml:"alap>nyelv"` // ENUM: hu, en, de, it, ro, sk, hr, fr, es, cz, pl
	AlapDevizanem     string  `xml:"alap>devizanem"`
	AlapDevizabank    string  `xml:"alap>devizabank,omitempty"`
	AlapDevizaarf     float64 `xml:"alap>devizaarf,omitempty"`
	AlapMegjegyzes    string  `xml:"alap>megjegyzes,omitempty"`
	AlapPenzforg      bool    `xml:"alap>penzforg"`
	AlapKata          bool    `xml:"alap>kata"`
	AlapEmail         string  `xml:"alap>email,omitempty"`
	AlapTeszt         bool    `xml:"alap>teszt"`

	VevoId                int    `xml:"vevo>id,omitempty"`
	VevoNev               string `xml:"vevo>nev"`
	VevoAzonosito         string `xml:"vevo>azonosito,omitempty"`
	VevoCimOrszag         string `xml:"vevo>cim>orszag"`
	VevoCimIrsz           string `xml:"vevo>cim>irsz"`
	VevoCimTelepules      string `xml:"vevo>cim>telepules"`
	VevoCimCim            string `xml:"vevo>cim>cim"`
	VevoPostacimOrszag    string `xml:"vevo>postacim>orszag,omitempty"`
	VevoPostacimIrsz      string `xml:"vevo>postacim>irsz,omitempty"`
	VevoPostacimTelepules string `xml:"vevo>postacim>telepules,omitempty"`
	VevoPostacimCim       string `xml:"vevo>postacim>cim,omitempty"`
	VevoEmail             string `xml:"vevo>email,omitempty"`
	VevoAdoszam           string `xml:"vevo>adoszam"`
	VevoAdoszameu         string `xml:"vevo>adoszameu,omitempty"`
	VevoLokacio           int    `xml:"vevo>lokacio"`

	VevoFokonyvVevo           string `xml:"vevo>fokonyv>vevo,omitempty"`
	VevoFokonyvVevoazon       string `xml:"vevo>fokonyv>vevoazon,omitempty"`
	VevoFokonyvDatum          Date   `xml:"vevo>fokonyv>datum,omitempty"`
	VevoFokonyvFolyamatostelj bool   `xml:"vevo>fokonyv>folyamatostelj,omitempty"`
	VevoFokonyvElszDatTol     Date   `xml:"vevo>fokonyv>elszDatTol,omitempty"`
	VevoFokonyvElszDatIg      Date   `xml:"vevo>fokonyv>elszDatIg,omitempty"`

	Tetelek []SzamlaTetel `xml:"tetelek>tetel"`

	OsszegekAfakulcsossz    []SzamlaAfakulcsossz `xml:"osszegek>afakulcsossz"`
	OsszegekTotalosszNetto  float64              `xml:"osszegek>totalossz>netto"`
	OsszegekTotalosszAfa    float64              `xml:"osszegek>totalossz>afa"`
	OsszegekTotalosszBrutto float64              `xml:"osszegek>totalossz>brutto"`

	Kifizetesek []SzamlaKifizetes `xml:"kifizetesek>kifizetes,omitempty"`

	Pdf []byte `xml:"pdf"`
}

// SzamlaTetel describes a line item within Szamla
type SzamlaTetel struct {
	Nev              string  `xml:"nev"`
	Azonosito        string  `xml:"azonosito,omitempty"`
	Mennyiseg        float64 `xml:"mennyiseg"`
	Mennyisegiegyseg string  `xml:"mennyisegiegyseg"`
	Nettoegysegar    float64 `xml:"nettoegysegar"`
	Afatipus         string  `xml:"afatipus,omitempty"` // ENUM: TAM, AAM, EU, EUK, MAA, F.AFA, K.AFA, ÁKK, TAHK, TEHK, EUT, EUKT
	Afakulcs         float64 `xml:"afakulcs"`           // schema states int, but field has decimals. Ex: 27.0
	Netto            float64 `xml:"netto"`
	Arresafaalap     float64 `xml:"arresafaalap,omitempty"`
	Afa              float64 `xml:"afa"`
	Brutto           float64 `xml:"brutto"`
	Megjegyzes       string  `xml:"megjegyzes,omitempty"`

	FokonyvArbevetel           string `xml:"fokonyv>arbevetel,omitempty"`
	FokonyvAfa                 string `xml:"fokonyv>afa,omitempty"`
	FokonyvGazdasagiesemeny    string `xml:"fokonyv>gazdasagiesemeny,omitempty"`
	FokonyvGazdasagiesemenyafa string `xml:"fokonyv>gazdasagiesemenyafa,omitempty"`
	FokonyvElszdattol          Date   `xml:"fokonyv>elszdattol,omitempty"`
	FokonyvElszdatig           Date   `xml:"fokonyv>elszdatig,omitempty"`
}

// SzamlaAfakulcsossz describes a summary for a VAT key within Szamla
type SzamlaAfakulcsossz struct {
	Afatipus string  `xml:"afatipus,omitempty"` // ENUM: TAM, AAM, EU, EUK, MAA, F.AFA, K.AFA, ÁKK, TAHK, TEHK, EUT, EUKT
	Afakulcs float64 `xml:"afakulcs"`           // schema states int, but field has decimals. Ex: 27.0
	Netto    float64 `xml:"netto"`
	Afa      float64 `xml:"afa"`
	Brutto   float64 `xml:"brutto"`
}

// SzamlaKifizetesek describes a payment item within Szamla
type SzamlaKifizetes struct {
	Datum          Date    `xml:"datum"`
	Jogcim         string  `xml:"jogcim"`
	Osszeg         float64 `xml:"osszeg"`
	Megjegyzes     string  `xml:"megjegyzes,omitempty"`
	Bankszamlaszam string  `xml:"bankszamlaszam,omitempty"`
	Banktranzid    int     `xml:"banktranzid,omitempty"`
	Devizaarf      float64 `xml:"devizaarf,omitempty"`
}

// Xmlszamladbkdel describes the request of deleting proforma invoices
type Xmlszamladbkdel struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlszamladbkdel xmlszamladbkdel"`

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`

	Szamlaszam   string `xml:"fejlec>szamlaszam,omitempty"`
	Rendelesszam string `xml:"fejlec>rendelesszam,omitempty"`
}

// Xmlszamladbkdelvalasz describes the response of deleting proforma invoices
type Xmlszamladbkdelvalasz struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlszamladbkdelvalasz xmlszamladbkdelvalasz"` // sic!

	Sikeres    bool   `xml:"sikeres"`
	Hibakod    int    `xml:"hibakod,omitempty"`
	Hibauzenet string `xml:"hibauzenet,omitempty"`
}

// Xmlnyugtacreate describes the request of creating receipts
type Xmlnyugtacreate struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlnyugtacreate xmlnyugtacreate"`

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`
	PdfLetoltes      bool   `xml:"beallitasok>pdfLetoltes"`

	HivasAzonosito string  `xml:"fejlec>hivasAzonosito"`
	Elotag         string  `xml:"fejlec>elotag"`
	Fizmod         string  `xml:"fejlec>fizmod"`
	Penznem        string  `xml:"fejlec>penznem"`
	Devizaarf      float64 `xml:"fejlec>devizaarf,omitempty"`
	Devizabank     string  `xml:"fejlec>devizabank,omitempty"`
	Megjegyzes     string  `xml:"fejlec>megjegyzes,omitempty"`
	PdfSablon      string  `xml:"fejlec>pdfSablon,omitempty"`
	FokonyvVevo    string  `xml:"fejlec>fokonyvVevo,omitempty"`

	Tetelek []XmlnyugtacreateTetel `xml:"tetelek>tetel"`

	Kifizetesek *[]XmlnyugtacreateKifizetes `xml:"kifizetesek>kifizetes,omitempty"`
}

// XmlnyugtacreateTetel is a line item of Xmlnyugtacreate
type XmlnyugtacreateTetel struct {
	Megnevezes       string  `xml:"megnevezes"`
	Azonosito        string  `xml:"azonosito,omitempty"`
	Mennyiseg        float64 `xml:"mennyiseg"`
	MennyisegiEgyseg string  `xml:"mennyisegiEgyseg"`
	NettoEgysegar    float64 `xml:"nettoEgysegar"`
	Afakulcs         string  `xml:"afakulcs"` // ENUM:  0, 5, 10, 27, AAM, TAM, EU, EUK, MAA, F.AFA, K.AFA, ÁKK,HO, EUE, EUFADE, EUFAD37, ATK, NAM, EAM, KBAUK, KBAET
	Netto            float64 `xml:"netto"`
	Afa              float64 `xml:"afa"`
	Brutto           float64 `xml:"brutto"`

	FokonyvArbevetel string `xml:"fokonyv>arbevetel,omitempty"`
	FokonyvAfa       string `xml:"fokonyv>afa,omitempty"`
}

// XmlnyugtacreateKifizetes is a payment line of Xmlnyugtacreate
type XmlnyugtacreateKifizetes struct {
	Fizetoeszkoz string  `xml:"fizetoeszkoz"`
	Osszeg       float64 `xml:"osszeg"`
	Leiras       string  `xml:"leiras,omitempty"`
}

// Xmlnyugtavalasz describes the response of creating receipts
type Xmlnyugtavalasz struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlnyugtavalasz xmlnyugtavalasz"`

	Sikeres    bool   `xml:"sikeres"`
	Hibakod    int    `xml:"hibakod,omitempty"`
	Hibauzenet string `xml:"hibauzenet,omitempty"`

	NyugtaPdf []byte `xml:"nyugtaPdf,omitempty"`

	Id                   int     `xml:"nyugta>alap>id"`
	HivasAzonosito       string  `xml:"nyugta>alap>hivasAzonosito"`
	Nyugtaszam           string  `xml:"nyugta>alap>nyugtaszam"`
	Tipus                string  `xml:"nyugta>alap>tipus"`
	Stornozott           bool    `xml:"nyugta>alap>stornozott"`
	StronozottNyugtaszam string  `xml:"nyugta>alap>stornozottNyugtaszam"`
	Kelt                 Date    `xml:"nyugta>alap>kelt"`
	Fizmod               string  `xml:"nyugta>alap>fizmod"`
	Penznem              string  `xml:"nyugta>alap>penznem"`
	Devizaarf            float64 `xml:"nyugta>alap>devizaarf,omitempty"`
	Devizabank           string  `xml:"nyugta>alap>devizabank,omitempty"`
	Megjegyzes           string  `xml:"nyugta>alap>megjegyzes,omitempty"`
	FokonyvVevo          string  `xml:"nyugta>alap>fokonyvVevo,omitempty"`
	Teszt                bool    `xml:"nyugta>alap>teszt"`

	Tetelek []XmlnyugtavalaszTetel `xml:"nyugta>tetelek>tetel"`

	Kifizetesek *[]XmlnyugtavalaszKifizetes `xml:"nyugta>kifizetesek>kifizetes,omitempty"`

	OsszegekAfakulcsossz    []XmlnyugtavalaszAfakulcsossz `xml:"nyugta>osszegek>afakulcsossz"`
	OsszegekTotalosszNetto  float64                       `xml:"nyugta>osszegek>totalossz>netto"`
	OsszegekTotalosszAfa    float64                       `xml:"nyugta>osszegek>totalossz>afa"`
	OsszegekTotalosszBrutto float64                       `xml:"nyugta>osszegek>totalossz>brutto"`
}

// XmlnyugtavalaszTetel is a line item of Xmlnyugtavalasz
type XmlnyugtavalaszTetel struct {
	Megnevezes       string  `xml:"megnevezes"`
	Azonosito        string  `xml:"azonosito,omitempty"`
	NettoEgysegar    float64 `xml:"nettoEgysegar"`
	Mennyiseg        float64 `xml:"mennyiseg"`
	MennyisegiEgyseg string  `xml:"mennyisegiEgyseg"`
	Netto            float64 `xml:"netto"`
	Afatipus         string  `xml:"afatipus"` // ENUM: TAM, AAM, EU, EUK, MAA, F.AFA, K.AFA, ÁKK, TAHK, TEHK, EUT, EUKT, HO, EUE, EUFADE, EUFADE37, ATK, NAM, EAM, KBAUK, KBAET
	Afakulcs         float64 `xml:"afakulcs"` // schema states int, but field has decimals. Ex: 27.0
	Afa              float64 `xml:"afa"`
	Brutto           float64 `xml:"brutto"`

	FokonyvArbevetel string `xml:"fokonyv>arbevetel,omitempty"`
	FokonyvAfa       string `xml:"fokonyv>afa,omitempty"`
}

// XmlnyugatavalaszKifizetes is a payment item of Xmlnyugtavalasz
type XmlnyugtavalaszKifizetes struct {
	Fizetoeszkoz string  `xml:"fizetoeszkoz"`
	Osszeg       float64 `xml:"osszeg"`
	Leiras       string  `xml:"leiras,omitempty"`
}

// XmlnyugtavalaszAfakulcsossz describes a summary for a VAT key within Xmlnyugtavalasz
type XmlnyugtavalaszAfakulcsossz struct {
	Afatipus string  `xml:"afatipus,omitempty"` // ENUM: TEHK, TAHK, TAM, AAM, EUT, EUKT, MAA, F.AFA, K.AFA, HO, EUE, EUFADE, EUFAD37, ATK, NAM, EAM, KBAUK, KBAET, 0, 5, 7, 18, 19, 20, 25, 27
	Afakulcs float64 `xml:"afakulcs"`           // schema states int, but field has decimals. Ex: 27.0
	Netto    float64 `xml:"netto"`
	Afa      float64 `xml:"afa"`
	Brutto   float64 `xml:"brutto"`
}

// Xmlnyugtast describes the request of reversing receipts
type Xmlnyugtast struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlnyugtast xmlnyugtast"`

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`
	PdfLetoltes      bool   `xml:"beallitasok>pdfLetoltes"`

	Nyugtaszam     string `xml:"fejlec>nyugtaszam"`
	PdfSablon      string `xml:"fejlec>pdfSablon,omitempty"`
	HivasAzonosito string `xml:"fejlec>hivasAzonosito"`
}

// Xmlnyugtaget describes the request of downloading receipts
type Xmlnyugtaget struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlnyugtaget xmlnyugtaget"`

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`
	PdfLetoltes      bool   `xml:"beallitasok>pdfLetoltes"`

	Nyugtaszam     string `xml:"fejlec>nyugtaszam"`
	PdfSablon      string `xml:"fejlec>pdfSablon,omitempty"`
	HivasAzonosito string `xml:"fejlec>hivasAzonosito"`
}

// Xmlnyugtasend describes the request of sending receipts via email
type Xmlnyugtasend struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlnyugtasend xmlnyugtasend"`

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`

	Nyugtaszam string `xml:"fejlec>nyugtaszam"`

	Email        string `xml:"emailKuldes>email,omitempty"`
	EmailReplyto string `xml:"emailKuldes>emailReplyto,omitempty"`
	EmailTargy   string `xml:"emailKuldes>emailTargy,omitempty"`
	EmailSzoveg  string `xml:"emailKuldes>emailSzoveg,omitempty"`
}

// Xmlnyugtasendvalasz describes the response of sending receipts via email
type Xmlnyugtasendvalasz struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlnyugtasendvalasz xmlnyugtasendvalasz"`

	Sikeres    bool   `xml:"sikeres"`
	Hibakod    int    `xml:"hibakod,omitempty"`
	Hibauzenet string `xml:"hibauzenet,omitempty"`
}

// Xmltaxpayer describes the request for querying NAV taxpayer details
type Xmltaxpayer struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmltaxpayer xmltaxpayer"`

	// Authentication (Username+Password or AgentKey required)

	Felhasznalo      string `xml:"beallitasok>felhasznalo,omitempty"`
	Jelszo           string `xml:"beallitasok>jelszo,omitempty"`
	SzamlaAgentKulcs string `xml:"beallitasok>szamlaagentkulcs,omitempty"`

	Torzsszam string `xml:"torzsszam"` // VAT no (first 8 chars)
}

// QueryTaxpayerResponse describes the response of querying NAV taxpayer details
// Note: This is a partial implementation based on NAV schema
type QueryTaxpayerResponse struct {
	XMLName Empty `xml:"http://schemas.nav.gov.hu/OSA/3.0/api QueryTaxpayerResponse"`

	ResultFuncCode  string `xml:"result>funcCode"` // ENUM: OK, ERROR
	ResultErrorCode string `xml:"result>errorCode,omitempty"`
	ResultMessage   string `xml:"result>message,omitempty"`

	Validity  bool   `xml:"taxpayerValidity,omitempty"`
	Name      string `xml:"taxpayerData>taxpayerName"`
	ShortName string `xml:"taxpayerData>taxpayerShortName,omitempty"`

	TaxpayerId string `xml:"taxpayerData>taxNumberDetail>taxpayerId"`
	VatCode    string `xml:"taxpayerData>taxNumberDetail>vatCode,omitempty"`
	CountyCode string `xml:"taxpayerData>taxNumberDetail>countyCode,omitempty"`

	VatGroupMembership string `xml:"taxpayerData>vatGroupMembership,omitempty"`

	AddressList []TaxpayerAddressItem `xml:"taxpayerData>taxpayerAddressList>taxpayerAddressItem"`
}

// TaxpayerAddressItem describes a single address within QueryTaxpayerResponse
type TaxpayerAddressItem struct {
	Type                string `xml:"taxpayerAddressType"` // ENUM: HQ, SITE, BRANCH
	Country             string `xml:"taxpayerAddress>countryCode"`
	PostalCode          string `xml:"taxpayerAddress>postalCode"`
	City                string `xml:"taxpayerAddress>city"`
	StreetName          string `xml:"taxpayerAddress>streetName"`
	PublicPlaceCategory string `xml:"taxpayerAddress>publicPlaceCategory"`
	Number              string `xml:"taxpayerAddress>number,omitempty"`
	Building            string `xml:"taxpayerAddress>building,omitempty"`
	Staircase           string `xml:"taxpayerAddress>staircase,omitempty"`
	Floor               string `xml:"taxpayerAddress>floor,omitempty"`
	Door                string `xml:"taxpayerAddress>door,omitempty"`
	LotNumber           string `xml:"taxpayerAddress>lotNumber,omitempty"`
}

// xmlcegmb describes the request for registering a managed user account
type Xmlcegmb struct {
	XMLName Empty `xml:"http://www.szamlazz.hu/xmlcegmb XmlCegMb"`

	LoginName        string `xml:"login>loginname"`        // API user
	Password         string `xml:"login>password"`         // API password
	SzamlaAgentKulcs string `xml:"login>szamlaagentkulcs"` // API token

	CompanyName      string `xml:"cegMb>cegcompanyname"`
	TaxNumber        string `xml:"cegMb>cegtaxnumber"`
	SzamlaszamElotag string `xml:"cegMb>cegszamlaszamelotag"`
	Irsz             string `xml:"cegMb>cegirsz"`
	City             string `xml:"cegMb>cegcity"`
	Addr             string `xml:"cegMb>cegaddr"`

	PostIrsz          string `xml:"cegMb>cegpostirsz,omitempty"`
	PostCity          string `xml:"cegMb>cegpostcity,omitempty"`
	PostAddr          string `xml:"cegMb>cegpostaddr,omitempty"`
	Bank              string `xml:"cegMb>cegbank"`
	BankAccount       string `xml:"cegMb>cegbankaccount"`
	Email             string `xml:"cegMb>cegemail"`
	EmailReplyTo      string `xml:"cegMb>cegemailreplyto"`
	CegPenzforgDatTol string `xml:"cegMb>cegpenzforgdattol,omitempty"`
	CegPenzforgDatIg  string `xml:"cegMb>cegpenzforgdatig,omitempty"`
	CegKataDatTol     string `xml:"cegMb>cegkatadattol,omitempty"`
	CegKataDatIg      string `xml:"cegMb>cegpenzkatadatig,omitempty"`

	UsrEmail      string `xml:"usrMb>usremail"`
	UsrPassword   string `xml:"usrMb>usrpassword"`
	UsrVezeteknev string `xml:"usrMb>usrvezeteknev"`
	UsrKeresztnev string `xml:"usrMb>usrkeresztnev"`
}
