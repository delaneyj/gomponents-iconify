package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outgoing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTOutgoing0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linejoin="round" d="M18 35a4 4 0 1 0-8 0a4 4 0 0 0 8 0Zm19 0a4 4 0 1 0-8 0a4 4 0 0 0 8 0Z"/><path stroke-linecap="round" d="M4 35h6m8 0h11m8 0h7"/><path stroke-linecap="round" stroke-linejoin="round" d="m38 19l6-6l-6-6M4 13h40"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTOutgoing0)"/>`),
		g.Group(children),
	)
}