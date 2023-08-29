package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LatinCross(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#8D65C5" d="M30 26a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20Z"/><path fill="#fff" d="M16 6a1 1 0 0 0-1 1v5h-5a1 1 0 1 0 0 2h5v11a1 1 0 1 0 2 0V14h5a1 1 0 1 0 0-2h-5V7a1 1 0 0 0-1-1Z"/></g>`),
		g.Group(children),
	)
}