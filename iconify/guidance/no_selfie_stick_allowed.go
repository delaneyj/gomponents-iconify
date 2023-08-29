package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoSelfieStickAllowed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m.5.5l23 23m-23 .4l.649-4.802A3 3 0 0 1 4.122 16.5H5.5l2.353 4.706A2.342 2.342 0 0 0 9.947 22.5c.364 0 .722-.084 1.024-.287c.709-.477 2.119-1.598 3.529-3.713M16 16l4-10.5m3.5-3.143V6.5h-.25L16.5 4.643V.5h.25l6.75 1.857ZM4.547 14.507s-2.434-.091-3.242-1.527c-.626-1.11-.258-2.527.821-3.166c1.08-.638 2.457-.253 3.082.857c.809 1.435-.326 3.638-.326 3.638l-.335.198Z"/>`),
		g.Group(children),
	)
}