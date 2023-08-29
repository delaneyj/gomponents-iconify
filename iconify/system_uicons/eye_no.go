package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeNo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.211 6.26C4.727 7.173 3.323 8.587 2 10.5c2.537 3.667 5.37 5.5 8.5 5.5c1.423 0 2.785-.38 4.085-1.137m1.588-1.14c.98-.843 1.922-1.916 2.827-3.223C16.463 6.833 13.63 5 10.5 5c-.83 0-1.64.13-2.429.387M4 4l13 13.071"/>`),
		g.Group(children),
	)
}