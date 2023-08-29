package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h18v18h1v2h-9v-2h1V10h-4v10h1v2H2v-2h1V2Zm2 18h3V10H5v10ZM5 8h14V4H5v4Zm14 2h-3v10h3V10Z"/>`),
		g.Group(children),
	)
}