package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AvTimer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M240 384h32v32h-32zM96 240h32v32H96zm288 0h32v32h-32z"/><path fill="currentColor" d="M414.392 97.608A222.332 222.332 0 0 0 272 32.567V32h-32v96h32V64.672C370.41 72.83 448 155.519 448 256c0 105.869-86.131 192-192 192S64 361.869 64 256a191.61 191.61 0 0 1 56.408-135.942l115.624 145.88l25.078-19.876L124.6 73.828l-12.606 10.59a224 224 0 1 0 302.4 13.19Z"/>`),
		g.Group(children),
	)
}