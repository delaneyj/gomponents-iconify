package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TargetOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M.5 7.5a7 7 0 1 0 14 0a7 7 0 0 0-14 0Z"/><path d="M3.5 7.5a4 4 0 1 0 8 0a4 4 0 0 0-8 0Z"/><path d="M6.5 7.5a1 1 0 1 0 2 0a1 1 0 0 0-2 0Z"/></g>`),
		g.Group(children),
	)
}