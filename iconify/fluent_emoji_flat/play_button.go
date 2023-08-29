package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M23.5 15.134a1 1 0 0 1 0 1.732l-12 6.928a1 1 0 0 1-1.5-.866V9.072a1 1 0 0 1 1.5-.866l12 6.928Z"/></g>`),
		g.Group(children),
	)
}