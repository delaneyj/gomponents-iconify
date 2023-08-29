package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BridgeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBridgeOne0"><path fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 19v14h8c0-6.627 5.373-12 12-12s12 5.373 12 12h8V19s-12.035-4-20-4s-20 4-20 4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBridgeOne0)"/>`),
		g.Group(children),
	)
}