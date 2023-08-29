package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M1.5 11A1.5 1.5 0 0 1 3 9.5h15a1.5 1.5 0 0 1 0 3H3A1.5 1.5 0 0 1 1.5 11Z" opacity=".2"/><path d="M.5 10a.5.5 0 0 1 .5-.5h18a.5.5 0 0 1 0 1H1a.5.5 0 0 1-.5-.5Z"/></g>`),
		g.Group(children),
	)
}