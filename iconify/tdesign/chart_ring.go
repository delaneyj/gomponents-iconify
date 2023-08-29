package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 3.055a9.001 9.001 0 1 0 6.631 15.966l-2.123-2.153A6 6 0 1 1 11 6.083V3.055Zm2 0v3.028A6.005 6.005 0 0 1 17.917 11h3.028A9.004 9.004 0 0 0 13 3.055ZM20.945 13h-3.028a5.973 5.973 0 0 1-1.004 2.445l2.13 2.16A8.957 8.957 0 0 0 20.945 13ZM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12Zm11-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}