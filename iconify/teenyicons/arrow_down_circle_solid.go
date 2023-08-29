package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowDownCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 0a7.5 7.5 0 1 1 0 15a7.5 7.5 0 0 1 0-15Zm2.707 8.5L7.5 11.207L4.793 8.5l.707-.707l1.5 1.5V4h1v5.293l1.5-1.5l.707.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}