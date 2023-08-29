package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiskTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="32" height="22" x="4" y="13" stroke="currentColor" stroke-width="4" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 13h10v22H4V13Zm32 6h8v10h-8"/><circle cx="30" cy="21" r="2" fill="currentColor"/><circle cx="30" cy="27" r="2" fill="currentColor"/></g>`),
		g.Group(children),
	)
}