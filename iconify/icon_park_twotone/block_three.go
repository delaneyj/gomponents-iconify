package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBlockThree0"><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M6 24h12v12H6V24Zm24 0h12v12H30V24ZM6 12h12v12H6V12Zm12 0h12v12H18V12Zm12 0h12v12H30V12Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBlockThree0)"/>`),
		g.Group(children),
	)
}