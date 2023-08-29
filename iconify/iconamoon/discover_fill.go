package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiscoverFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm13-3l-3.693.739a2 2 0 0 0-1.568 1.568L9 15l3.693-.739a2 2 0 0 0 1.569-1.568L15 9Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}