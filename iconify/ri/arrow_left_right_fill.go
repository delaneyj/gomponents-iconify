package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftRightFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 16v-4l5 5l-5 5v-4H4v-2h12ZM8 2v3.999L20 6v2H8v4L3 7l5-5Z"/>`),
		g.Group(children),
	)
}