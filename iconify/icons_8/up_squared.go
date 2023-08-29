package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpSquared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5H5zm2 2h18v18H7V7zm9 4.094l-.72.687l-6 6l1.44 1.44L16 13.937l5.28 5.28l1.44-1.437l-6-6l-.72-.686z"/>`),
		g.Group(children),
	)
}