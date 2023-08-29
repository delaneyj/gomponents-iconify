package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M736 897q-119 0-203.5-84.5T448 609q0-15 2-32h302q10-11 14.5-64t-4.5-64l-68-125q23-3 42-3q119 0 203.5 84.5T1024 609t-84.5 203.5T736 897zm-42-384H169q-10-10-10-25t10-26L409 11q9-11 22.5-11T454 11l240 451q10 11 9.5 26t-9.5 25zm-310 96q0 81 34 151.5T512 881v112q0 13-9.5 22.5T480 1025H32q-13 0-22.5-9.5T0 993V545q0-13 9.5-22.5T32 513h65q4 53 14 64h275q-2 24-2 32z"/>`),
		g.Group(children),
	)
}