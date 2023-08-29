package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickToRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M176 56v120H16v160h160v120h38.623l200-199.8L214.635 56Zm32 361.384V304H48v-96h160V94.639l161.373 161.535ZM463.998 56h32v400h-32z"/>`),
		g.Group(children),
	)
}