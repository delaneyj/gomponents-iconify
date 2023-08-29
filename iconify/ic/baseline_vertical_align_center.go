package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaselineVerticalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19h3v4h2v-4h3l-4-4l-4 4zm8-14h-3V1h-2v4H8l4 4l4-4zM4 11v2h16v-2H4z"/>`),
		g.Group(children),
	)
}