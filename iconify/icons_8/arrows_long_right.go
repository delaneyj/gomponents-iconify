package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsLongRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m21.188 9.28l-1.407 1.44L24.063 15H4v2h20.063l-4.282 4.28l1.407 1.44l6-6l.72-.72l-.72-.72l-6-6z"/>`),
		g.Group(children),
	)
}