package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.595 4.577c3.223-1.176 7.025-.61 9.65 1.63c2.982 2.543 3.601 6.523 1.636 9.66m-1.908 2.109C15.186 20.166 11.083 20.642 7.7 19L3 20l1.3-3.9C2.071 12.804 2.806 8.589 5.98 6.043M3 3l18 18"/>`),
		g.Group(children),
	)
}