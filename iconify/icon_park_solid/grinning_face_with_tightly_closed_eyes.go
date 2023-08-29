package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrinningFaceWithTightlyClosedEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGrinningFaceWithTightlyClosedEyes0"><g fill="none" stroke-linejoin="round" stroke-width="4"><path fill="#fff" stroke="#fff" d="M24 44c11.046 0 20-8.954 20-20S35.046 4 24 4S4 12.954 4 24s8.954 20 20 20Z"/><path stroke="#000" stroke-linecap="round" d="M17 31s2 4 7 4s7-4 7-4M16 16l3 3l-3 3m16-6l-3 3l3 3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGrinningFaceWithTightlyClosedEyes0)"/>`),
		g.Group(children),
	)
}