package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleLeftAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.43 238.43 0 0 1 256 496Zm0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48Z"/><path fill="currentColor" d="M284.687 379.313L161.373 256l123.314-123.313l22.626 22.626L206.628 256l100.685 100.687l-22.626 22.626z"/>`),
		g.Group(children),
	)
}