package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaggageDelay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBaggageDelay0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M36 26V14a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v24a2 2 0 0 0 2 2h17M16 12v28m12-28v17m0-17V6a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v6"/><path fill="#555" d="M35 44a9 9 0 1 0 0-18a9 9 0 0 0 0 18Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 32v4h4"/><path stroke-linecap="round" d="M13 40v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBaggageDelay0)"/>`),
		g.Group(children),
	)
}