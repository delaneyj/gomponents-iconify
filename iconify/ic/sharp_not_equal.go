package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpNotEqual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 9.998H5v-2h14zm0 6H5v-2h14z"/><path fill="currentColor" d="m14.08 4.605l1.84.79l-6 14l-1.84-.79z"/>`),
		g.Group(children),
	)
}