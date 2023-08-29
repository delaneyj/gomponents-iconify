package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FunctionArgumentRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 10a8 8 0 0 1 14.93-4h2.238A10.002 10.002 0 0 0 10 0C4.477 0 0 4.477 0 10s4.477 10 10 10c4.1 0 7.625-2.468 9.168-6H16.93A8 8 0 0 1 2 10Zm15.938-1H12V5l-6 5l6 5v-4h8V9h-2.062Z"/>`),
		g.Group(children),
	)
}