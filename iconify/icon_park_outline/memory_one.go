package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 28h38v14H5zM5 6h38v14H5z"/><rect width="4" height="4" x="11" y="11" fill="currentColor" rx="2"/><rect width="4" height="4" x="11" y="33" fill="currentColor" rx="2"/><rect width="4" height="4" x="19" y="11" fill="currentColor" rx="2"/><rect width="4" height="4" x="19" y="33" fill="currentColor" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M31 13h4m-4 22h4"/></g>`),
		g.Group(children),
	)
}