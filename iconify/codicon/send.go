package codicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Send(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="m1 1.91l.78-.41L15 7.449v.95L1.78 14.33L1 13.91L2.583 8L1 1.91ZM3.612 8.5L2.33 13.13L13.5 7.9L2.33 2.839l1.282 4.6L9 7.5v1H3.612Z"/>`),
		g.Group(children),
	)
}