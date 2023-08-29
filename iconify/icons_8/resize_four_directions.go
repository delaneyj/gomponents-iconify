package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeFourDirections(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 4v9h2V7.437L14.563 16L6 24.563V19H4v9h9v-2H7.437L16 17.437L24.563 26H19v2h9v-9h-2v5.563L17.437 16L26 7.437V13h2V4h-9v2h5.563L16 14.563L7.437 6H13V4H4z"/>`),
		g.Group(children),
	)
}