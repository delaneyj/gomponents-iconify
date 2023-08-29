package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M5 4h6.5a3 3 0 0 1 3 3v12"/><path d="M18 16.5L14.5 20L11 16.5"/></g>`),
		g.Group(children),
	)
}