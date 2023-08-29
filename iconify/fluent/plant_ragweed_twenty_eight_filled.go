package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlantRagweedTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 2a4 4 0 0 0-3.947 4.653a3.001 3.001 0 0 0-1.037 5.097A3 3 0 0 0 11 17h2.25v5.69l-4.47-4.47a.75.75 0 0 0-1.06 1.06l5.53 5.53v1.44a.75.75 0 0 0 1.5 0v-1.44l5.53-5.53a.75.75 0 1 0-1.06-1.06l-4.47 4.47V17H17a3 3 0 0 0 1.984-5.25a3.001 3.001 0 0 0-1.038-5.097A4 4 0 0 0 14 2Z"/>`),
		g.Group(children),
	)
}