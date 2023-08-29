package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 384H576v64h258q34 1 43 10l137 94q10 10 10 24t-10 24l-137 94q-10 10-45 10H576v256q0 26-18.5 45t-45.5 19t-45.5-19t-18.5-45V704H64q-26 0-45-19T0 640V512q0-26 19-45t45-19h384v-64H192q-35 0-45-10L10 280Q0 270 0 256t10-24l137-94q9-10 43-10h258V64q0-26 18.5-45T512 0t45.5 19T576 64v64h384q26 0 45 18.5t19 45.5v128q0 27-19 45.5T960 384z"/>`),
		g.Group(children),
	)
}