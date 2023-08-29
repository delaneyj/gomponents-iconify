package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dcu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.817 38.592V5.507h10.789a76.442 76.442 0 0 1 9.284.448c4.625.548 14.293 5.45 14.293 18.038a17.663 17.663 0 0 1-14.293 17.94a65.763 65.763 0 0 1-9.34.563l-10.733-.074L17.55 31.689v1.071h5.126c4.214 0 6.995-3.217 7.103-8.768s-2.889-8.845-6.478-8.845H17.55V28.14ZM18.494 15l8.833-8.834m2.556 16.384l8.547-8.547M10.323 42.299l9.39-9.391"/>`),
		g.Group(children),
	)
}