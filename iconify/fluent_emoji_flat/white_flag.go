package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WhiteFlag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#9B9B9B" d="M3.028 6a1 1 0 0 0-1 1v18a1 1 0 0 0 1 1h26a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-26Z"/><path fill="#fff" d="M3.028 7h26v18h-26V7Z"/></g>`),
		g.Group(children),
	)
}