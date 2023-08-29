package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSketch0"><g fill="#555" stroke="#fff" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" rx="3"/><path stroke-linecap="round" d="M18.6 16h10.8l3.6 4.706L24 32l-9-11.294L18.6 16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSketch0)"/>`),
		g.Group(children),
	)
}