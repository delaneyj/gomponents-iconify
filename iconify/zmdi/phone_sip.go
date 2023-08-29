package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneSip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M299 0v107h-22V0h22zm-43 43v64h-64V85h43V64h-43V0h64v21h-43v22h43zm64-43h64v64h-43v43h-21V0zm43 43V21h-22v22h22zm0 224q8 0 14.5 6t6.5 15v75q0 8-6.5 14.5T363 384q-99 0-182.5-48.5t-132-132T0 21q0-8 6.5-14.5T21 0h75q9 0 15 6.5t6 14.5q0 40 12 76q4 13-5 22l-47 47q48 93 141 141l47-47q9-9 22-5q36 12 76 12z"/>`),
		g.Group(children),
	)
}