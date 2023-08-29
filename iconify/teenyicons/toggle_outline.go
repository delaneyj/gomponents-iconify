package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M3.5 2.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2Z"/><path d="M11.5.5h-8a3 3 0 0 0 0 6h8a3 3 0 1 0 0-6Zm0 12a1 1 0 1 1 0-2a1 1 0 0 1 0 2Z"/><path d="M3.5 14.5h8a3 3 0 1 0 0-6h-8a3 3 0 0 0 0 6Z"/></g>`),
		g.Group(children),
	)
}