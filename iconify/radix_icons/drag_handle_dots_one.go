package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DragHandleDotsOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<g fill="currentColor"><circle cx="4.5" cy="2.5" r=".6"/><circle cx="4.5" cy="4.5" r=".6"/><circle cx="4.5" cy="6.499" r=".6"/><circle cx="4.5" cy="8.499" r=".6"/><circle cx="4.5" cy="10.498" r=".6"/><circle cx="4.5" cy="12.498" r=".6"/><circle cx="6.5" cy="2.5" r=".6"/><circle cx="6.5" cy="4.5" r=".6"/><circle cx="6.5" cy="6.499" r=".6"/><circle cx="6.5" cy="8.499" r=".6"/><circle cx="6.5" cy="10.498" r=".6"/><circle cx="6.5" cy="12.498" r=".6"/><circle cx="8.499" cy="2.5" r=".6"/><circle cx="8.499" cy="4.5" r=".6"/><circle cx="8.499" cy="6.499" r=".6"/><circle cx="8.499" cy="8.499" r=".6"/><circle cx="8.499" cy="10.498" r=".6"/><circle cx="8.499" cy="12.498" r=".6"/><circle cx="10.499" cy="2.5" r=".6"/><circle cx="10.499" cy="4.5" r=".6"/><circle cx="10.499" cy="6.499" r=".6"/><circle cx="10.499" cy="8.499" r=".6"/><circle cx="10.499" cy="10.498" r=".6"/><circle cx="10.499" cy="12.498" r=".6"/></g>`),
		g.Group(children),
	)
}