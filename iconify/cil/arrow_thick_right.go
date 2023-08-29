package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowThickRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M254.3 496h-38.275l-.008-143.937H16v-192h200.007L216 16.048l38.688.035l239.948 240.235ZM48 320.063h200.015l.007 137.006l201.342-200.8L248 54.672l.008 137.391H48Z"/>`),
		g.Group(children),
	)
}