package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alazan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.12 20.836a17.957 17.957 0 0 1-11.085 16.591a17.957 17.957 0 0 1-19.576-3.892a17.957 17.957 0 0 1-3.892-19.576a17.957 17.957 0 0 1 16.59-11.085"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.234 20.634A21.498 21.498 0 0 1 33.76 43.156A21.498 21.498 0 0 1 4.844 14.24A21.498 21.498 0 0 1 27.366 2.766m-1.341 17.69l11.128 6.703m-9.99-4.426l.11-15.947M42.25 11.87l-2.718-1.602l-2.699 1.565l.684-3.08l-2.323-2.083l3.142-.301l1.263-2.853l1.257 2.895l3.103.32l-2.364 2.09l.654 3.049Z"/>`),
		g.Group(children),
	)
}