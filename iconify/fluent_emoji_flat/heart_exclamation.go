package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartExclamation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#F92F60" d="M15.998 5.21s1.71-2.68 3.938-3.08c5.51-.994 7.788 3.943 6.864 7.616C25.16 16.28 15.998 22 15.998 22S6.836 16.28 5.196 9.746c-.913-3.673 1.356-8.61 6.865-7.617c2.229.402 3.937 3.081 3.937 3.081ZM19 27a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`),
		g.Group(children),
	)
}