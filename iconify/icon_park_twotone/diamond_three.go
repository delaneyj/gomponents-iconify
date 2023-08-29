package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDiamondThree0"><path fill="#555" stroke="#fff" stroke-width="4" d="M5.414 22.586L22.586 5.414a2 2 0 0 1 2.828 0l17.172 17.172a2 2 0 0 1 0 2.828L25.414 42.586a2 2 0 0 1-2.828 0L5.414 25.414a2 2 0 0 1 0-2.828Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDiamondThree0)"/>`),
		g.Group(children),
	)
}