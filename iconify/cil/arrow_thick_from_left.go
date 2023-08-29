package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickFromLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M296.007 56h-38.632v120h-144v160h144v120H296l200-199.8Zm-6.632 361.384V304h-144v-96h144V94.639l161.37 161.535ZM17.375 56h32v400h-32z"/>`),
		g.Group(children),
	)
}