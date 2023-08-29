package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsCNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM25.95 14.192a10 10 0 0 1 5.121 2.737A2 2 0 1 0 33.9 14.1a14 14 0 1 0 0 19.799a2 2 0 0 0-2.829-2.829a10.001 10.001 0 1 1-5.12-16.879Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsCNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}