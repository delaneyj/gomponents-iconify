package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaVerticalOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 22q.875-1.875 1.438-4.5T5.5 12q0-2.875-.563-5.5T3.5 2h17.025q-.875 1.875-1.45 4.5T18.5 12q0 2.875.575 5.5t1.45 4.5H3.5Zm14.15-2q-.575-1.95-.863-3.963T16.5 12q0-2.025.288-4.038T17.65 4H6.375q.575 1.95.85 3.963T7.5 12q0 2.025-.275 4.038T6.375 20H17.65ZM12 12Z"/>`),
		g.Group(children),
	)
}