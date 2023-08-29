package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlasWaterLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h14M5 4h-.5M5 4l1 8m13-8h.5M19 4l-1 8M6 12l.781 6.248A2 2 0 0 0 8.766 20h6.468a2 2 0 0 0 1.985-1.752L18 12M6 12l2.072-.69a2 2 0 0 1 1.742.233l1.077.717a2 2 0 0 0 2.218 0l1.077-.717a2 2 0 0 1 1.742-.234L18 12"/>`),
		g.Group(children),
	)
}