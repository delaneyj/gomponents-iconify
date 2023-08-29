package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm11.836-4.452a1 1 0 1 0-1.672-1.096l-3.489 5.323a4 4 0 1 0 3.55-1.769l1.611-2.458ZM10 14a2 2 0 1 1 4 0a2 2 0 0 1-4 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}