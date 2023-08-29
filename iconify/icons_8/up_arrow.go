package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16 4.094l-.72.687l-8.5 8.5l1.44 1.44L15 7.936V28h2V7.937l6.78 6.782l1.44-1.44l-8.5-8.5l-.72-.686z"/>`),
		g.Group(children),
	)
}