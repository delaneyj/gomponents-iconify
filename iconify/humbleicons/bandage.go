package humbleicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bandage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v.001M14 12v.001M12 14v.001M10 12v.001M19 12l-7 7a4.95 4.95 0 1 1-7-7l7-7a4.95 4.95 0 0 1 7 7Z"/>`),
		g.Group(children),
	)
}