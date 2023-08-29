package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UTurnLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSUTurnLeft0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M14 13h19c6.075 0 11 4.925 11 11v0c0 6.075-4.925 11-11 11H4"/><path stroke-linecap="round" stroke-linejoin="round" d="m9 40l-5-5l5-5"/><circle cx="9" cy="13" r="5" fill="#fff"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSUTurnLeft0)"/>`),
		g.Group(children),
	)
}