package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M11.038 9.722c0-.966.784-1.75 1.75-1.75h7.474a1.75 1.75 0 0 1 1.52 2.619l-7.253 12.695a1.75 1.75 0 0 1-3.039-1.736l5.757-10.078h-4.459a1.75 1.75 0 0 1-1.75-1.75Z"/></g>`),
		g.Group(children),
	)
}