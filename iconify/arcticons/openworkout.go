package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Openworkout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.62 39.693L43.5 27.84L24.01 8.306L4.5 27.775l11.89 11.916Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.217 32.106L24.012 16.868l-15.234 15.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.934 36.374l-10.92-10.945l-10.958 10.933"/>`),
		g.Group(children),
	)
}