package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toptal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.227 10.038L10.188 0l-2.04 2.04l3.773 3.769l-8.155 8.153L13.807 24l2.039-2.039l-3.772-3.771l8.16-8.152h-.007zM8.301 14.269l6.066-6.063l1.223 1.223l-6.064 6.113l-1.223-1.26l-.002-.013z"/>`),
		g.Group(children),
	)
}