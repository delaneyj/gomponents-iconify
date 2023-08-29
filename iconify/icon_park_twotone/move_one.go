package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTMoveOne0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="m8 6l35 19l-19 2l-10.005 17L8 6Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTMoveOne0)"/>`),
		g.Group(children),
	)
}