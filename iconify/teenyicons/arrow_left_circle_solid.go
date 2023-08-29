package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 7.5a7.5 7.5 0 1 1-15 0a7.5 7.5 0 0 1 15 0Zm-8.5 2.707L3.793 7.5L6.5 4.793l.707.707l-1.5 1.5H11v1H5.707l1.5 1.5l-.707.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}