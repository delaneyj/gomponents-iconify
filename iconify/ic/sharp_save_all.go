package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSaveAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 6h11l4 4v11H6V6zm2 2h7v3H8V8zm5.5 11a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5z" clip-rule="evenodd"/><path fill="currentColor" d="M2 2h12v2H4v10H2V2z"/>`),
		g.Group(children),
	)
}