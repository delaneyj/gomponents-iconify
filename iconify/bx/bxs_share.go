package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M11 6.914V2.586L6.293 7.293l-3.774 3.774l3.841 3.201L11 18.135V13.9c8.146-.614 11 4.1 11 4.1c0-2.937-.242-5.985-2.551-8.293C16.765 7.022 12.878 6.832 11 6.914z" fill="currentColor"/>`),
		g.Group(children),
	)
}