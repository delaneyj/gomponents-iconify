package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Store(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M363 21v43H21V21h342zm21 214h-21v128h-43V235h-85v128H21V235H0v-43L21 85h342l21 107v43zm-192 85v-85H64v85h128z"/>`),
		g.Group(children),
	)
}