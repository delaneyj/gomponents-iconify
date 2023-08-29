package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpsideDownFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 14.5v.01H9v-.01zm6 0v.01H15v-.01z"/><path stroke-linecap="round" d="M8.535 10A3.998 3.998 0 0 1 12 8c1.48 0 2.773.804 3.465 2"/></g>`),
		g.Group(children),
	)
}