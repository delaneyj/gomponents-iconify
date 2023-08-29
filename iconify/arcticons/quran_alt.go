package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuranAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5s-1.31 2.776-6.025 4.103s-5.495 1.939-5.495 5.728m-1.998.245H24M24 36H11.065M24 43.5s-2.617-2.805-5.63-2.834s-6.042-1.55-6.042-4.623m1.695.014V14.775m6.547-.047v21.27M24 4.5s1.311 2.776 6.026 4.103s5.495 1.939 5.495 5.728m1.997.245H24M24 36h12.936M24 43.5s2.618-2.805 5.63-2.834s6.043-1.55 6.043-4.623m-1.695.014V14.775m-6.548-.047v21.27"/>`),
		g.Group(children),
	)
}