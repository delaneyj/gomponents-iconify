package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 501q0 136-67 251T751 934t-251 67t-251-67T67 752T0 501t67-251T249 68T500 1t251 67t182 182t67 251zM547 733l71-71l-161-161l161-161l-71-71l-233 232z"/>`),
		g.Group(children),
	)
}