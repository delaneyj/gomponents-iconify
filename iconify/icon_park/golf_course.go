package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfCourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><ellipse cx="24" cy="34" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="20" ry="10"/><circle cx="32" cy="34" r="2" fill="#2F88FF" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path fill="#2F88FF" d="M24 9L13 4V14L24 9Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 34V14M13 14V4L24 9L13 14Z"/></g>`),
		g.Group(children),
	)
}