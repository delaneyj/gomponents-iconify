package fad

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cutter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M70.526 8.975L215.553 155l-90.48 89.973L68.86 187.58L70.526 8.975zM85.246 47l1.018 134.104L125.74 223l55.346-55.78l-54.06-54.22l11.872-12.917L85.246 47z"/>`),
		g.Group(children),
	)
}