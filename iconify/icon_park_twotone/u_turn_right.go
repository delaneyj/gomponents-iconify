package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UTurnRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTUTurnRight0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M34 13H15C8.925 13 4 17.925 4 24v0c0 6.075 4.925 11 11 11h29"/><path stroke-linecap="round" stroke-linejoin="round" d="m39 30l5 5l-5 5"/><circle cx="39" cy="13" r="5" fill="#555"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTUTurnRight0)"/>`),
		g.Group(children),
	)
}