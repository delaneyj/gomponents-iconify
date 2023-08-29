package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 501q0 136-67 251T751 934t-251 67t-251-67T67 752T0 501t67-251T249 68T500 1t251 67t182 182t67 251zm-117 0q0-158-112.5-270.5T500 118q-126 0-226 74l2 2q-48 35-83 83l-2-2q-74 100-74 226q0 158 112.5 270.5T500 884q126 0 226-74l-2-2q48-35 83-83l2 2q74-100 74-226zM680 701q0 25-17.5 42.5T620 761q-26 0-42-18L458 623q-18-18-18-62V241q0-25 17.5-42.5T500 181t42.5 17.5T560 241v315l102 103q18 16 18 42z"/>`),
		g.Group(children),
	)
}