package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoltCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M425.706 86.294A240 240 0 0 0 86.294 425.706A240 240 0 0 0 425.706 86.294ZM256 464c-114.691 0-208-93.309-208-208S141.309 48 256 48s208 93.309 208 208s-93.309 208-208 208Z"/><path fill="currentColor" d="M372.529 120H180.468L121.8 296h76.047l-36.4 104h77.179L384 254.627V208h-65.249ZM352 240v1.373L225.373 368h-18.821l36.4-104H166.2l37.333-112h111.938l-53.778 88Z"/>`),
		g.Group(children),
	)
}