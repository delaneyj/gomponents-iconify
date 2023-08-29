package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.545 17.454v3.273a3.273 3.273 0 1 1-3.272-3.273Zm8.728-8.727V12A3.273 3.273 0 0 1 12 15.273H5.455a3.273 3.273 0 0 1 0-6.546zM24 0v3.273a3.273 3.273 0 0 1-3.273 3.272H7.637a3.273 3.273 0 0 1 0-6.545Z"/>`),
		g.Group(children),
	)
}