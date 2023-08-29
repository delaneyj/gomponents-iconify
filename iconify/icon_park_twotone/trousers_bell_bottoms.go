package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrousersBellBottoms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTrousersBellBottoms0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="m16 12l2-8h12l2 8v12l4 17l-12 3l-12-3l4-17V12Z"/><path d="M24 44V16M12 41l12 3l12-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTrousersBellBottoms0)"/>`),
		g.Group(children),
	)
}