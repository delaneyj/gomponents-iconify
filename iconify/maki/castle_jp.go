package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleJp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M12.5 12.75a1 1 0 0 1-1-1v-4.5H10a1 1 0 0 1-1-1v-2H6v2a1 1 0 0 1-1 1H3.5v4.5a1 1 0 0 1-2 0v-5.5a1 1 0 0 1 1-1H4v-2a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v2h1.5a1 1 0 0 1 1 1v5.5a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}