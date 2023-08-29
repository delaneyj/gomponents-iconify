package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UTurnUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUTurnUp0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M13 14v19c0 6.075 4.925 11 11 11v0c6.075 0 11-4.925 11-11V4"/><path stroke-linecap="round" stroke-linejoin="round" d="m30 9l5-5l5 5"/><circle cx="13" cy="9" r="5" fill="#555" transform="rotate(-90 13 9)"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUTurnUp0)"/>`),
		g.Group(children),
	)
}