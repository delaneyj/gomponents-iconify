package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBlockFive0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 30h12v12H6V30Zm12-12h12v12H18V18ZM30 6h12v12H30V6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBlockFive0)"/>`),
		g.Group(children),
	)
}