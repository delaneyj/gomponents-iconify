package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DivingBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path stroke-miterlimit="2" d="m4 40l2.5 1.351c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0c1.6-.864 3.4-.864 5 0c1.6.865 3.4.865 5 0L44 40M14 11V4m17 7V4M21 8V4m17 4V4"/><path d="M12 6h8m9 0h8"/><rect width="8" height="24" x="10" y="11" rx="4"/><rect width="8" height="24" x="27" y="11" rx="4"/><path d="M27 15a4 4 0 0 1 8 0v4h-8v-4Zm-17 0a4 4 0 0 1 8 0v4h-8v-4Z"/></g>`),
		g.Group(children),
	)
}