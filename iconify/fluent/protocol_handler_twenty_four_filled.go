package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtocolHandlerTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m13.482 17.732l-.775-.775l2.482-2.482a3.5 3.5 0 0 0 0-4.95l-2.482-2.482l.775-.776a2.5 2.5 0 0 1 3.535 0l3.965 3.965a2.5 2.5 0 0 1 0 3.535l-3.965 3.965a2.5 2.5 0 0 1-3.535 0Zm1-3.965a2.5 2.5 0 0 0 0-3.535l-3.965-3.965a2.5 2.5 0 0 0-3.535 0l-3.965 3.965a2.5 2.5 0 0 0 0 3.535l3.965 3.965a2.5 2.5 0 0 0 3.535 0l3.965-3.965Z"/>`),
		g.Group(children),
	)
}