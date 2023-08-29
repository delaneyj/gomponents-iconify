package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerStandby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 3h-2v10h2V3m4.83 2.17l-1.42 1.42A6.944 6.944 0 0 1 19 12a7 7 0 0 1-7 7A6.995 6.995 0 0 1 7.58 6.58L6.17 5.17a9.001 9.001 0 0 0-1.03 12.69c3.22 3.78 8.9 4.24 12.69 1.02A9.003 9.003 0 0 0 21 12c0-2.63-1.16-5.13-3.17-6.83Z"/>`),
		g.Group(children),
	)
}