package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Glasswire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.51 23.52c12.6 0 23.24 9.93 33.63 9.93a44.19 44.19 0 0 0 7.54-.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.12 41.24c10.66-2.81 16.86-23 31.73-27.59"/>`),
		g.Group(children),
	)
}