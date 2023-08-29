package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Appointment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><circle cx="24" cy="11" r="7" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M4 41C4 32.1634 12.0589 25 22 25"/><circle cx="34" cy="34" r="9" fill="#2F88FF" stroke="#000"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M33 31V35H37"/></g>`),
		g.Group(children),
	)
}