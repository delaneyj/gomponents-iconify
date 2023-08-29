package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDiamondTwo0"><path fill="#555" stroke="#fff" stroke-width="4" d="m8.923 22.788l13.486-17.7a2 2 0 0 1 3.182 0l13.486 17.7a2 2 0 0 1 0 2.424l-13.486 17.7a2 2 0 0 1-3.182 0l-13.486-17.7a2 2 0 0 1 0-2.424Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDiamondTwo0)"/>`),
		g.Group(children),
	)
}