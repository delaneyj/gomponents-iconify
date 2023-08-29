package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m6.007 13.418l2-12l.986.164l-2 12l-.986-.164Zm-.8-8.918l-3 3l3 3l-.707.707L.793 7.5L4.5 3.793l.707.707Zm5.293-.707L14.207 7.5L10.5 11.207l-.707-.707l3-3l-3-3l.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}