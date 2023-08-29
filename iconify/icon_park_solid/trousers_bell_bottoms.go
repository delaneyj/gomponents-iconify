package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrousersBellBottoms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSTrousersBellBottoms0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="m16 12l2-8h12l2 8v12l4 17l-12 3l-12-3l4-17V12Z"/><path stroke="#000" d="M24 44V16"/><path stroke="#fff" d="m12 41l12 3l12-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSTrousersBellBottoms0)"/>`),
		g.Group(children),
	)
}