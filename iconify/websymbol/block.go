package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Block(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 501q0 136-67 251T751 934t-251 67t-251-67T67 752T0 501t67-251T249 68T500 1t251 67t182 182t67 251zm-117 0q0-158-112.5-270.5T500 118q-126 0-226 74l535 535q74-100 74-226zM726 810L191 275q-74 100-74 226q0 158 112.5 270.5T500 884q126 0 226-74z"/>`),
		g.Group(children),
	)
}