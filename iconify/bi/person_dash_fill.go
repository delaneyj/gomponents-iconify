package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonDashFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 7.5a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 0 1h-4a.5.5 0 0 1-.5-.5z"/><path d="M1 14s-1 0-1-1s1-4 6-4s6 3 6 4s-1 1-1 1H1zm5-6a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/></g>`),
		g.Group(children),
	)
}