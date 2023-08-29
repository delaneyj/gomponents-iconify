package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsBookNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM14 5a5 5 0 0 0-5 5v32a1 1 0 0 0 1 1h28v-2H11v-1a3 3 0 0 1 3-3h24a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H14Zm9 16v5h2v-5h5v-2h-5v-5h-2v5h-5v2h5Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsBookNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}