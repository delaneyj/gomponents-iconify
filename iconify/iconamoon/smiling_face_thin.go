package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M15.465 14A3.998 3.998 0 0 1 12 16a3.998 3.998 0 0 1-3.465-2M14 10l1-1l1 1m-6 0L9 9l-1 1"/></g>`),
		g.Group(children),
	)
}