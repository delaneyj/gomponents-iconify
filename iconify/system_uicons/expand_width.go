package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M1.5 2.5v16.021M19.5 2.5v16.021m-5.993-4.006l4-4l-4-4.015m-6 8.015l-4-4l4-4.015m9.993 4h-14"/>`),
		g.Group(children),
	)
}