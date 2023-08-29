package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsLongUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4.094l-.72.718l-6 6l1.44 1.407L15 7.936V28h2V7.937l4.28 4.282l1.44-1.408l-6-6l-.72-.718z"/>`),
		g.Group(children),
	)
}