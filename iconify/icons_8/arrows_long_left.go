package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsLongLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m10.813 9.28l-6 6l-.72.72l.72.72l6 6l1.406-1.44L7.936 17H28v-2H7.937l4.282-4.28l-1.408-1.44z"/>`),
		g.Group(children),
	)
}