package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Launch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m733 1017l-93-249h-23q-8 23-24 66t-23 62H448v64q0 27-19 45.5t-45.5 18.5t-45-19t-18.5-45v-64H197q-6-19-22-62t-24-66h-23l-93 249Q0 999 0 960V768q0-27 51-77.5t77-50.5q0-136 32.5-267T252 142T384 0q73 42 132 142t91.5 231T640 640q26 0 77 50.5t51 77.5v192q0 39-35 57zM384 256q-53 0-90.5 37.5T256 384t37.5 90.5T384 512t90.5-37.5T512 384t-37.5-90.5T384 256z"/>`),
		g.Group(children),
	)
}