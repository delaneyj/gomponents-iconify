package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockSix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBlockSix0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 18h12v12H18V18ZM6 18h12v12H6V18Zm24 0h12v12H30V18Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBlockSix0)"/>`),
		g.Group(children),
	)
}