package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPaintedPlate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTHandPaintedPlate0"><g fill="none"><path stroke="#fff" stroke-width="4" d="M18 10h24a2 2 0 0 1 2 2v26a2 2 0 0 1-2 2H18"/><path fill="#555" stroke="#fff" stroke-width="4" d="M4 12a2 2 0 0 1 2-2h12v30H6a2 2 0 0 1-2-2V12Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M11 17v4m26-5L25 33"/><rect width="4" height="4" x="9" y="25" fill="#fff" rx="2"/><rect width="4" height="4" x="9" y="31" fill="#fff" rx="2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTHandPaintedPlate0)"/>`),
		g.Group(children),
	)
}