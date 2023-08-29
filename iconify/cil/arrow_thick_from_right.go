package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickFromRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M254.625 56h-38.632L16 256.2L216 456h38.623V336h144V176h-144Zm112 152v96h-144v113.384l-161.37-161.21l161.37-161.535V208Zm96-152h32v400h-32z"/>`),
		g.Group(children),
	)
}