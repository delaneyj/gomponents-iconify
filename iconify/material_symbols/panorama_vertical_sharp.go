package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaVerticalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 22q.875-1.875 1.438-4.5T5.5 12q0-2.875-.563-5.5T3.5 2h17.025q-.875 1.875-1.45 4.5T18.5 12q0 2.875.575 5.5t1.45 4.5H3.5Z"/>`),
		g.Group(children),
	)
}