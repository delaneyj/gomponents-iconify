package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsGNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM30.667 16.546a2 2 0 1 0 2.666-2.981a14 14 0 1 0 0 20.87a2 2 0 0 0 .667-1.49V24a2 2 0 0 0-2-2h-8a2 2 0 1 0 0 4h6v6a10 10 0 1 1 .667-15.454Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsGNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}