package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftVisualStudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17 8.5l-4.75 3.82L17 16V8.5M4.7 18.4L2 16.7v-9l3-1l4.3 3.33L18 2l4 2.5V20l-5 2l-7.66-7.34L4.7 18.4M5 14l1.86-1.72L5 10.5V14Z"/>`),
		g.Group(children),
	)
}