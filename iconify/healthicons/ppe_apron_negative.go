package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PpeApronNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPpeApronNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM19.428 7.027A6 6 0 0 0 14 13v12a2 2 0 0 0 2 2h2v12l2.838.946a10 10 0 0 0 6.324 0L30 39V27h2a2 2 0 0 0 2-2V13a6 6 0 0 0-5.428-5.973a5.001 5.001 0 0 1-9.144 0ZM19 25v-2h10v2H19Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPpeApronNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}