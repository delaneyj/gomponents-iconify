package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerLeftUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 19h9v-2h-7V6.828l3.95 3.95l1.414-1.414L11 3L4.636 9.364l1.414 1.414L10 6.828V19Z"/>`),
		g.Group(children),
	)
}