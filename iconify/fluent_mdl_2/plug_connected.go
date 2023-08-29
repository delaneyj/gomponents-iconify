package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugConnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1216q0-89 34-171t97-146l227-226l633 633l-226 227q-63 63-145 97t-172 34q-73 0-141-22t-127-67l-327 326l-90-90l326-327q-44-58-66-126t-23-142zm448 320q64 0 122-24t104-70l136-136l-452-452l-136 136q-45 45-69 103t-25 123q0 66 25 124t68 102t102 69t125 25zm871-1100q44 58 66 126t23 142q0 89-34 171t-97 146l-227 226l-633-633l226-227q63-63 145-97t172-34q73 0 141 22t127 67l327-326l90 90l-326 327zm-133 494q45-45 69-103t25-123q0-66-25-124t-69-101t-102-69t-124-26q-64 0-122 24t-104 70L854 614l452 452l136-136z"/>`),
		g.Group(children),
	)
}