package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleDownAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.43 238.43 0 0 1 256 496Zm0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48Z"/><path fill="currentColor" d="M256 342.627L132.687 219.313l22.626-22.626L256 297.372l100.687-100.685l22.626 22.626L256 342.627z"/>`),
		g.Group(children),
	)
}