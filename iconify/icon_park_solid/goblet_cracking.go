package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GobletCracking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGobletCracking0"><g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke="#fff" d="M33 44H13m10-16v16"/><path fill="#fff" stroke="#fff" d="M35 16c0 6.5-5.373 12-12 12s-12-5.373-12-12c0-6.5 4-12 4-12h16s4 5.5 4 12Z"/><path stroke="#000" d="m23 4l-2 6l4 1l-2 6"/><path stroke="#fff" d="M15 4h16"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGobletCracking0)"/>`),
		g.Group(children),
	)
}