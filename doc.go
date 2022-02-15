/*
Package szamlazzhu is a Go language binding for the szamlazz.hu Számla Agent API

The NewUserAgent and NewTokenAgent functions should be used to create an
Agent struct initialized with számlázz.hu credentials. Methods of the struct
can be used to invoke API endpoints afterwards.

	a := szamlazzhu.NewUserAgent("user@example.com", "P4ssword")
	receipt, err := a.GenerateReceipt(szamlazzhu.Xmlnyugtacreate{
		Elotag: "BNY",
		Fizmod: "készpénz",
		Tetelek: []szamlazzhu.XmlnyugtacreateTetel{
			{
				Megnevezes:       "Eladott dolog",
				NettoEgysegar:    100,
				Mennyiseg:        5,
				MennyisegiEgyseg: "db",
				Netto:            500,
				Afakulcs:         "27",
				Afa:              135,
				Brutto:           635,
			},
		},
	})
	if err != nil {
		fmt.Printf("Something went wrong: %v\n", err)
	}else{
		fmt.Printf("Receipt issued, number: %s\n", receipt.Nyugtaszam)
	}

The agent struct should be kept as-long as you wish to use the same credentials,
as it stores session-specific identifier.

Each method is a direct call to an API endpoint, taking the parameters as
described by the API documentation. The fields Felhasználónév, Jelszó, Token and
VálaszVerzió will be set by the Agent, and should not need to be used directly.

All errors - including stacktraces, HTTP headers, XML-embedded failure codes and
timeouts - are collected and returned in convenient Go format, irrespective of
the format the REST API returns them.

When err == nil, it is safe to assume the business-logic of the call has
happened successfully. When err != nil, an error has happened meaning that the
business-logic of the call may or may not happened. The client needs to examine
the error, and decide how to proceed, including any retry attempts if
applicable.

In addition to this documentation, you may refer to https://docs.szamlazz.hu
for additional information about the API itself.

*/
package szamlazzhu
