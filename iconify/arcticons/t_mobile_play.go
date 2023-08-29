package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TMobilePlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.54 4.5H10.658v39h8.21V28.105h6.672c6.518 0 11.802-5.284 11.802-11.802S32.058 4.5 25.54 4.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.849 16.303L18.84 9.369v13.868l12.009-6.934z"/>`),
		g.Group(children),
	)
}