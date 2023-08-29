package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsFNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM18 10a2 2 0 0 0-2 2v24a2 2 0 1 0 4 0V26h10a2 2 0 1 0 0-4H20v-8h10a2 2 0 1 0 0-4H18Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsFNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}