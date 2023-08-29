package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerRightUpLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 19H5v-2h7V6.828l-3.95 3.95l-1.414-1.414L13 3l6.364 6.364l-1.414 1.414L14 6.828V19Z"/>`),
		g.Group(children),
	)
}