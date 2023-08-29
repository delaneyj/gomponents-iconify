package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSketch0"><g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" rx="3"/><path fill="#000" stroke="#000" stroke-linecap="round" d="M18.6 16h10.8l3.6 4.706L24 32l-9-11.294L18.6 16Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSketch0)"/>`),
		g.Group(children),
	)
}