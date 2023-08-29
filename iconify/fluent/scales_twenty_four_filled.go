package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScalesTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.75 3a.75.75 0 0 0 0 1.5h1.042l-2.737 6.717A.75.75 0 0 0 2 11.5a3.5 3.5 0 1 0 7 0a.75.75 0 0 0-.055-.283L6.208 4.5h5.042v12H7.253a2.25 2.25 0 0 0 0 4.5h9.497a2.25 2.25 0 0 0 0-4.5h-4v-12h5.042l-2.737 6.717A.75.75 0 0 0 15 11.5a3.5 3.5 0 1 0 7 0a.75.75 0 0 0-.055-.283L19.208 4.5h1.042a.75.75 0 0 0 0-1.5H3.75ZM5.5 6.738l1.635 4.012h-3.27L5.5 6.738Zm11.365 4.012L18.5 6.738l1.635 4.012h-3.27Z"/>`),
		g.Group(children),
	)
}