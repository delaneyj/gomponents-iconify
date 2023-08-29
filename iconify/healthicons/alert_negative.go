package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlertNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsAlertNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm21 7a3 3 0 1 1 6 0v24a3 3 0 1 1-6 0V7Zm3 31a3 3 0 1 0 0 6a3 3 0 0 0 0-6Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsAlertNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}