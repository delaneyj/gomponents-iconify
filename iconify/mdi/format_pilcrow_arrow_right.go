package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatPilcrowArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m21 18l-4-4v3H5v2h12v3M9 10v5h2V4h2v11h2V4h2V2H9a4 4 0 0 0-4 4a4 4 0 0 0 4 4Z"/>`),
		g.Group(children),
	)
}