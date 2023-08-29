package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SilverwareVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.1 13.34L3.91 9.16a4.008 4.008 0 0 1 0-5.66l7.02 7l-2.83 2.84m5.31-.34l6.88 6.88l-1.41 1.41L12 14.41l-6.88 6.88l-1.41-1.41l9.65-9.66l-.2-.22a1.98 1.98 0 0 1 0-2.81l4.34-4.37l.93.92L15.19 7l.96.94l3.24-3.25l.92.92l-3.25 3.24l.94.96l3.26-3.25l.92.94l-4.37 4.34c-.78.78-2.04.78-2.81 0l-.22-.2L13.41 13Z"/>`),
		g.Group(children),
	)
}