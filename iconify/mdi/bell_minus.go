package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BellMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 19v-2h-8v2h8M12 2c-1.1 0-2 .9-2 2v.29C7.12 5.14 5 7.82 5 11v6l-2 2v1h9.35a5.93 5.93 0 0 1-.35-2c0-3.31 2.69-6 6-6c.34 0 .67.03 1 .09V11c0-3.18-2.12-5.86-5-6.71V4a2 2 0 0 0-2-2m-2 19a2 2 0 0 0 3.65 1.13c-.32-.34-.6-.72-.84-1.13H10Z"/>`),
		g.Group(children),
	)
}