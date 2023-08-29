package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Viber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 4.5a1.88 1.88 0 0 0-1.92 1.88v30.07A1.88 1.88 0 0 0 9 38.33h1.88v5.17l10.34-5.17H39a1.88 1.88 0 0 0 1.88-1.88V6.38A1.88 1.88 0 0 0 39 4.5Zm4.7 6.58h5.6a1 1 0 0 1 .94.94a10.06 10.06 0 0 0 .63 3.76a1.85 1.85 0 0 1-.63 2.35l-2 2a17.94 17.94 0 0 0 3.38 4.6a18.94 18.94 0 0 0 4.61 3.39l2-2a1.85 1.85 0 0 1 2.35-.63a10.06 10.06 0 0 0 3.76.63a1 1 0 0 1 .94.94v5.64a1 1 0 0 1-.94.94A22.53 22.53 0 0 1 12.72 12a1 1 0 0 1 .94-.92Z"/>`),
		g.Group(children),
	)
}