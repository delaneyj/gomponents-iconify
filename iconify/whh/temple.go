package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Temple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 726v234q0 27-18.5 45.5T832 1024H576V896q0-27-18.5-45.5t-45-18.5t-45.5 18.5t-19 45.5v128H192q-26 0-45-18.5T128 960V726q-56-47-92-91.5T0 572q0-53 32-60q11 4 67.5 26t92.5 35V448q-57-35-92.5-73T64 316q0-60 32-60q14 5 57 23.5t76.5 29.5t58.5 11q27 0 54.5-18.5t54-52.5t44-62t40.5-68q-33-18-33-55q0-27 19-45.5T512.5 0t45 18.5T576 64q0 37-33 55q22 40 40 68t44.5 62t54 52.5T736 320q25 0 58.5-11t76.5-29.5t57-23.5q32 0 32 60q0 21-35.5 59T832 448v125q36-13 92.5-35t67.5-26q32 7 32 60q0 18-36 62.5T896 726z"/>`),
		g.Group(children),
	)
}