package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassesThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTGlassesThree0"><g fill="none" stroke="#fff" stroke-width="4"><circle cx="12" cy="33" r="7" fill="#555"/><path stroke-linecap="round" d="M29 33c0-3.314-1.5-6-5-6s-5 2.686-5 6"/><circle cx="36" cy="33" r="7" fill="#555"/><path stroke-linecap="round" d="M14 8h-2a6 6 0 0 0-6 6v8M34 8h2a6 6 0 0 1 6 6v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTGlassesThree0)"/>`),
		g.Group(children),
	)
}