package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.5 6.5c.5-2.5 4.343-2.657 6-1c1.603 1.603 1.5 4.334 0 6l-6 6l-6-6a4.243 4.243 0 0 1 0-6c1.55-1.55 5.5-1.5 6 1z"/>`),
		g.Group(children),
	)
}