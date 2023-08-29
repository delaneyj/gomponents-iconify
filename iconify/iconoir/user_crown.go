package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserCrown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M5 20v-1a7 7 0 0 1 10-6.326M21 22l1-6l-3.5 1.8L17 16l-1.5 1.8L12 16l1 6h8Z"/></g>`),
		g.Group(children),
	)
}