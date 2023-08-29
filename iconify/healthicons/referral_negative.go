package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReferralNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsReferralNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM34 31c1.934 0 3.5-1.566 3.5-3.5S35.934 24 34 24a3.499 3.499 0 0 0-3.5 3.5c0 1.934 1.566 3.5 3.5 3.5Zm-8 6.364C26 34.462 31.33 33 34 33s8 1.462 8 4.364V42H17c-1.271 0-2.298-.487-2.997-1.273C13.32 39.959 13 38.963 13 38v-3.586l-3.293 3.293l-1.414-1.414l5-5l.707-.707l.707.707l5 5l-1.414 1.414L15 34.414V38c0 .537.18 1.041.497 1.398c.301.339.774.602 1.503.602h9v-2.636ZM13 6a2 2 0 0 0-2 2v1H6v19h22V9h-5V8a2 2 0 0 0-2-2h-8Zm-5 5h3v2H8v-2Zm0 4h3v1a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1h3v11h-5v-6h-8v6H8V15Zm18-2h-3v-2h3v2Zm-7 9v4h-4v-4h4ZM16 8v3h-3v2h3v3h2v-3h3v-2h-3V8h-2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsReferralNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}