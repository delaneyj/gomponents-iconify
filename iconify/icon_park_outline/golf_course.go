package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfCourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><ellipse cx="24" cy="34" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="20" ry="10"/><circle cx="32" cy="34" r="2" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path d="M24 9L13 4v10l11-5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 34V14m0 0V4l11 5l-11 5Z"/></g>`),
		g.Group(children),
	)
}