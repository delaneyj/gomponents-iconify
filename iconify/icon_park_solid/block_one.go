package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBlockOne0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 30h12v12H6V30ZM30 6h12v12H30V6Zm0 12h12v12H30V18Zm-12 0h12v12H18V18ZM6 18h12v12H6V18Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBlockOne0)"/>`),
		g.Group(children),
	)
}