package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="m6.5 5.5l1 1l1-1a1.414 1.414 0 0 1 2 2l-3 3l-3-3a1.414 1.414 0 0 1 2-2Z"/><path d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/></g>`),
		g.Group(children),
	)
}