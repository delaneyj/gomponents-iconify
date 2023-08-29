package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesklampSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m6.854 1.146l1.884 1.885a4.551 4.551 0 0 1 4.98.98c.27.27.27.708 0 .979L7.99 10.717c-.27.27-.71.27-.98 0a4.551 4.551 0 0 1-.979-4.979l-.984-.984L2.166 8.46L7.707 14H13v1H2v-1h4.293L1.146 8.854a.5.5 0 0 1-.04-.66L4.333 4.04l-.188-.187a1.914 1.914 0 1 1 2.708-2.708Z"/>`),
		g.Group(children),
	)
}