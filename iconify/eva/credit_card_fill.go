package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCreditCardFill0"><g id="evaCreditCardFill1"><path id="evaCreditCardFill2" fill="currentColor" d="M19 5H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3Zm-8 10H7a1 1 0 0 1 0-2h4a1 1 0 0 1 0 2Zm6 0h-2a1 1 0 0 1 0-2h2a1 1 0 0 1 0 2Zm3-6H4V8a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1Z"/></g></g>`),
		g.Group(children),
	)
}