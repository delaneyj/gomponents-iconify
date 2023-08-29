package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HospitalBed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 17V39"/><path d="M42 25L42 39"/><path d="M26 15H38"/><path d="M11 22H17"/><path d="M6 28L42 28"/><path d="M6 34L42 34"/><path d="M32 9V21"/></g>`),
		g.Group(children),
	)
}