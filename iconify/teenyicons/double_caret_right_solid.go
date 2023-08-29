package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleCaretRightSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.474.293L9.234 7.5l-7.76 7.207l-.68-.733L7.764 7.5L.793 1.026l.68-.733Zm5.5 0l7.76 7.207l-7.76 7.207l-.68-.733l6.97-6.474l-6.971-6.474l.68-.733Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}