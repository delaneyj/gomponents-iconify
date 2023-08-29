package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockEight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBlockEight0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 6h12v12H6V6Zm12 0h12v12H18V6Zm0 12h12v12H18V18Zm0 12h12v12H18V30ZM30 6h12v12H30V6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBlockEight0)"/>`),
		g.Group(children),
	)
}