package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteLargeSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#9B9B9B" d="M4.028 1.5a2.5 2.5 0 0 0-2.5 2.5v24a2.5 2.5 0 0 0 2.5 2.5h24a2.5 2.5 0 0 0 2.5-2.5V4a2.5 2.5 0 0 0-2.5-2.5h-24Z"/><path fill="#fff" d="M2.528 4a1.5 1.5 0 0 1 1.5-1.5h24a1.5 1.5 0 0 1 1.5 1.5v24a1.5 1.5 0 0 1-1.5 1.5h-24a1.5 1.5 0 0 1-1.5-1.5V4Z"/></g>`),
		g.Group(children),
	)
}