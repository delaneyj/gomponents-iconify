package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneVerticalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 1v4H8l4 4l4-4h-3V1zM4 11h16v2H4zm4 8h3v4h2v-4h3l-4-4z"/>`),
		g.Group(children),
	)
}