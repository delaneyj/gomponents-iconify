package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14 5v2h9.563L7 23.563V14H5v13h13v-2H8.437L25 8.437V18h2V5H14z"/>`),
		g.Group(children),
	)
}