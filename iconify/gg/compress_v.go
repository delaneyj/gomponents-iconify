package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompressV(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.034 7.356L15.5 4.854l1.424 1.404l-4.913 4.985L7.025 6.33L8.43 4.905l2.604 2.566l.05-6.627l2 .015l-.05 6.497Zm2.529 11.176l1.412-1.416l-4.957-4.943l-4.943 4.957l1.417 1.412l2.584-2.592l.026 7.207l2-.008l-.026-7.096l2.487 2.479Z"/>`),
		g.Group(children),
	)
}