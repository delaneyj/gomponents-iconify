package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zend(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.28 9.81S11.28 7 14.09 7h8.41s0 2.81-2.79 2.81h-8.43m0 3.6s0-2.81 2.81-2.81h4.22s0 2.81-2.81 2.81h-4.22m0 3.59s0-2.81 2.81-2.81h1.41s0 2.81-2.82 2.81h-1.4m-.82-2.2V17H1.58L7.3 9.21H2.4V7h9.26l-5.7 7.8h4.5Z"/>`),
		g.Group(children),
	)
}