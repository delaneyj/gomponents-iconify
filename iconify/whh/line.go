package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Line(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 192h-64q-7 0-15-2L190 881q2 8 2 15v64q0 26-18.5 45t-45.5 19H64q-26 0-45-19T0 960v-64q0-27 19-45.5T64 832h64q7 0 15 2l691-691q-2-8-2-15V64q0-26 18.5-45T896 0h64q26 0 45 19t19 45v64q0 27-19 45.5T960 192z"/>`),
		g.Group(children),
	)
}