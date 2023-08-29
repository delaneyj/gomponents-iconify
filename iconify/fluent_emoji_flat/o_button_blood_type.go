package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OButtonBloodType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#F92F60" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M10.031 12.977a5.977 5.977 0 1 1 11.953 0v6.046a5.977 5.977 0 1 1-11.953 0v-6.046ZM16.008 10a2.977 2.977 0 0 0-2.977 2.977v6.046a2.977 2.977 0 0 0 5.953 0v-6.046A2.977 2.977 0 0 0 16.008 10Z"/></g>`),
		g.Group(children),
	)
}