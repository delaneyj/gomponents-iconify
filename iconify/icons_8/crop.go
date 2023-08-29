package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 4v4H4v2h4v14h14v4h2v-4h4v-2H11.437L22 11.437V21h2V9.437l3.72-3.718l-1.44-1.44L22.563 8H11v2h9.563L10 20.563V4H8z"/>`),
		g.Group(children),
	)
}