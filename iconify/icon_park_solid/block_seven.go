package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockSeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBlockSeven0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M18 18h12v12H18V18Zm12 0h12v12H30V18ZM6 18h12v12H6V18Zm12 12h12v12H18V30Zm0-24h12v12H18V6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBlockSeven0)"/>`),
		g.Group(children),
	)
}