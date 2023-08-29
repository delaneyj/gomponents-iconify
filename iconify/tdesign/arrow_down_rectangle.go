package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2V2Zm2 2v16h16V4H4Zm9 2.5v7.586l3-3l1.414 1.414L12 17.914L6.586 12.5L8 11.086l3 3V6.5h2Z"/>`),
		g.Group(children),
	)
}