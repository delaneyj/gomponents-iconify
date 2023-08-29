package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetStraightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M104 56v40H48V56a8 8 0 0 1 8-8h40a8 8 0 0 1 8 8Zm96-8h-40a8 8 0 0 0-8 8v40h56V56a8 8 0 0 0-8-8Z" opacity=".2"/><path d="M200 40h-40a16 16 0 0 0-16 16v88a16 16 0 0 1-32 0V56a16 16 0 0 0-16-16H56a16 16 0 0 0-16 16v88a88 88 0 0 0 88 88h.67c48.15-.36 87.33-40.29 87.33-89V56a16 16 0 0 0-16-16Zm0 16v32h-40V56ZM96 56v32H56V56Zm32.55 160A72 72 0 0 1 56 144v-40h40v40a32 32 0 0 0 64 0v-40h40v39c0 40-32 72.71-71.45 73Z"/></g>`),
		g.Group(children),
	)
}