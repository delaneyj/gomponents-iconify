package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tallglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 1024H64q-26 0-45-19T0 960V32Q0 19 9.5 9.5T32 0h448q13 0 22.5 9.5T512 32v928q0 27-19 45.5t-45 18.5zM138 810l80 22q13 3 16-10l22-80q3-13-10-16l-80-22q-13-3-16 10l-22 80q-1 5 1.5 10t8.5 6zM68 570l65 65q11 11 21 0l65-65q11-10 0-21l-65-65q-10-11-21 0l-65 65q-11 11 0 21zM448 64H64v256h384V64zm0 602l-22-80q-1-6-6-8.5t-10-1.5l-80 22q-13 3-10 16l22 80q1 6 6 8.5t10 1.5l80-22q13-3 10-16z"/>`),
		g.Group(children),
	)
}