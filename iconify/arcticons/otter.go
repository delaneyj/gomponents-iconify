package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Otter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.076 23.98a8.28 8.28 0 1 1-.004-.24m4.375-7.698l.09 15.916m6.654-15.916l.09 16.096m6.475-11.06l.09 5.935m6.474-8.003l.18 10.16"/>`),
		g.Group(children),
	)
}