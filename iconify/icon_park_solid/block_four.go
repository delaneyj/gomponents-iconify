package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBlockFour0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 12h12v12H6V12Zm12 0h12v12H18V12Zm12 0h12v12H30V12ZM18 24h12v12H18V24Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBlockFour0)"/>`),
		g.Group(children),
	)
}