package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.929 14.07l-3.43 3.43l-6-6a4.243 4.243 0 0 1 0-6a2.96 2.96 0 0 1 .567-.438m2.453-.605c1.388.034 2.706.668 2.98 2.043c.5-2.5 4.344-2.657 6-1c1.604 1.603 1.5 4.334 0 6l-.937.937M4 4l13 13.071"/>`),
		g.Group(children),
	)
}