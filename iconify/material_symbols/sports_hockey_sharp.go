package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsHockeySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-4h2v4H2Zm3 0v-4h4l.85-1.95l1.6 3.5l-1.1 2.45H5Zm15 0v-4h2v4h-2Zm-1 0h-5.35L6.35 4H9.7L12 9.2L14.3 4h3.35l-4.05 8.85L15 16h4v4Z"/>`),
		g.Group(children),
	)
}