package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPersonNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm24 26c5.523 0 10-4.477 10-10S29.523 6 24 6s-10 4.477-10 10s4.477 10 10 10Zm0 2c-6.008 0-18 3.035-18 8v6h36v-6c0-4.965-11.992-8-18-8Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPersonNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}