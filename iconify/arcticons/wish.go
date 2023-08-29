package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.917 28.754a5.84 5.84 0 0 0 5.909 6.983a8.448 8.448 0 0 0 8.057-6.983"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.883 28.754a5.84 5.84 0 0 0 5.908 6.983a8.448 8.448 0 0 0 8.058-6.983m0 0l.651-9.063m-14.617 9.063l.651-5.868m-14.617 5.868l1.246-11.811L4.5 12.263"/>`),
		g.Group(children),
	)
}