package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M13 10V35"/><path stroke-linejoin="round" d="M35 10V35"/><path d="M8 18L40 18"/><path stroke-linejoin="round" d="M24 10V18"/><path fill="#2F88FF" stroke-linejoin="round" d="M39 10H9.00001L5 4C5 4 16.0708 5 24 5C31.9292 5 43 4 43 4L39 10Z"/><rect width="6" height="9" x="10" y="35" fill="#2F88FF" stroke-linejoin="round"/><rect width="6" height="9" x="32" y="35" fill="#2F88FF" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}