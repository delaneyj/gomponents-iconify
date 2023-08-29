package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditCardOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M16 3.5H4a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3Zm-13 3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-7Z" clip-rule="evenodd"/><path d="M2.5 6.5h15a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-15A.5.5 0 0 1 2 8V7a.5.5 0 0 1 .5-.5Z"/><path fill-rule="evenodd" d="M6 9.5H5a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1Zm-1 2v-1h1v1H5Z" clip-rule="evenodd"/><path d="M1.293 2.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/></g>`),
		g.Group(children),
	)
}