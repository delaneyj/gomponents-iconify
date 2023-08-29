package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M31 15.25A6.25 6.25 0 0 0 24.75 9h-14.5A6.25 6.25 0 0 0 4 15.25v17.5A6.25 6.25 0 0 0 10.25 39h14.5A6.25 6.25 0 0 0 31 32.75v-17.5Zm2 14.282l6.961 5.436C41.603 36.25 44 35.08 44 32.998V15.003c0-2.083-2.397-3.252-4.039-1.97L33 18.468v11.063Z"/>`),
		g.Group(children),
	)
}