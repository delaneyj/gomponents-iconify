package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 21h6V3H3v18Zm1-2h4V5H4v14Zm1-2h2V7H5v10Zm10 4h6V3h-6v18Zm1-2h4V5h-4v14Zm1-2h2V7h-2v10Z"/>`),
		g.Group(children),
	)
}