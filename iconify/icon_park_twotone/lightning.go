package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lightning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTLightning0"><path fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4" d="M19 4h18L26 18h15L17 44l5-19H8L19 4Z"/></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTLightning0)"/>`),
		g.Group(children),
	)
}