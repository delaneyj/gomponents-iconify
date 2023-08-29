package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Width(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.507 14.515l4-4l-4-4.015m-10 8.015l-4-4l4-4.015m13.993 4h-18"/>`),
		g.Group(children),
	)
}