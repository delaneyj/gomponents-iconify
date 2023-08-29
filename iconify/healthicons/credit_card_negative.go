package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCreditCardNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM4 13a3 3 0 0 1 3-3h34a3 3 0 0 1 3 3v4H4v-4Zm0 10h40v12a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3V23Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCreditCardNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}