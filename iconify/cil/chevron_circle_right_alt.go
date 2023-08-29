package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronCircleRightAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 496A240 240 0 0 1 86.294 86.294a240 240 0 0 1 339.412 339.412A238.43 238.43 0 0 1 256 496Zm0-448C141.309 48 48 141.309 48 256s93.309 208 208 208s208-93.309 208-208S370.691 48 256 48Z"/><path fill="currentColor" d="m227.313 379.313l-22.626-22.626L305.372 256L204.687 155.313l22.626-22.626L350.627 256L227.313 379.313z"/>`),
		g.Group(children),
	)
}