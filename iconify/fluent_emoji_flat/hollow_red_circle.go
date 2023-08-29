package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HollowRedCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#F92F60" d="M16 7.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17ZM2 16C2 8.268 8.268 2 16 2s14 6.268 14 14s-6.268 14-14 14S2 23.732 2 16Z"/>`),
		g.Group(children),
	)
}