package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 2h18v5h1v2h-1v11h1v2h-9v-2h1v-5a2 2 0 1 0-4 0v5h1v2H2v-2h1V9H2V7h1V2Zm2 7v11h3v-5a4 4 0 0 1 8 0v5h3V9H5Zm14-2V4H5v3h14Z"/>`),
		g.Group(children),
	)
}