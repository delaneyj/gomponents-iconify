package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaCreditCardOutline0"><g id="evaCreditCardOutline1"><g id="evaCreditCardOutline2" fill="currentColor"><path d="M19 5H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3ZM4 8a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v1H4Zm16 8a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-5h16Z"/><path d="M7 15h4a1 1 0 0 0 0-2H7a1 1 0 0 0 0 2Zm8 0h2a1 1 0 0 0 0-2h-2a1 1 0 0 0 0 2Z"/></g></g></g>`),
		g.Group(children),
	)
}