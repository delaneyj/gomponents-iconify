package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v4h16V4H4Zm16 6h-4v10h4V10Zm-6 10V10h-4v10h4Zm-6 0V10H4v10h4Z"/>`),
		g.Group(children),
	)
}