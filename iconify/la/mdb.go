package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mdb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m1.08 9l-1 12h2.008l.557-6.695L5.322 21h1.356l2.677-6.695L9.912 21h2.008l-1-12H9.322L6 17.309L2.678 9H1.08zM14 9v12h2c3.302 0 6-2.698 6-6s-2.698-6-6-6h-2zm10 0v12.012h4.494A3.521 3.521 0 0 0 32 17.506c0-1.27-.723-2.342-1.744-2.957c.436-.584.744-1.27.744-2.049c0-1.921-1.579-3.5-3.5-3.5H24zm-8 2c2.22 0 4 1.78 4 4c0 2.22-1.78 4-4 4v-8zm10 0h1.5c.84 0 1.5.66 1.5 1.5s-.66 1.5-1.5 1.5H26v-3zm0 5h2.494A1.49 1.49 0 0 1 30 17.506a1.49 1.49 0 0 1-1.506 1.506H26V16z"/>`),
		g.Group(children),
	)
}