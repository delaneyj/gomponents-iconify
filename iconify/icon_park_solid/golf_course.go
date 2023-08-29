package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GolfCourse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSGolfCourse0"><g fill="none"><ellipse cx="24" cy="34" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="20" ry="10"/><circle cx="32" cy="34" r="2" fill="#fff" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path fill="#fff" d="M24 9L13 4v10l11-5Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 34V14m0 0V4l11 5l-11 5Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSGolfCourse0)"/>`),
		g.Group(children),
	)
}