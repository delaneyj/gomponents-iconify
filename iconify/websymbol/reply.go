package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M501 0v319q110 15 203.5 45t159 68T982 514.5t86 90.5t57 90.5t35 83.5t17.5 69t6.5 47l1 17q-6-11-19-29.5t-67.5-63T974 743t-196-52q-77-12-162-12q-55 0-115 5v316L0 501z"/>`),
		g.Group(children),
	)
}