package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.432 238.432 0 0 1 256 496Zm0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48Z"/><path fill="currentColor" d="M356.686 315.313L256 214.628L155.314 315.313l-22.628-22.626L256 169.373l123.314 123.314l-22.628 22.626z"/>`),
		g.Group(children),
	)
}