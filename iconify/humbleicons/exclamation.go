package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M13.253 5.98L12 13.5l-1.253-7.52a1.27 1.27 0 1 1 2.506 0Z"/><circle cx="12" cy="19" r="1" stroke-width="2"/></g>`),
		g.Group(children),
	)
}