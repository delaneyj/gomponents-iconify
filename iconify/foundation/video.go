package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M66.272 61.337v10.565c0 2.497-2.043 4.535-4.539 4.535H14.54c-2.496 0-4.54-2.038-4.54-4.535V28.097c0-2.496 2.043-4.535 4.54-4.535h47.193c2.496 0 4.539 2.038 4.539 4.535v10.432v-.146L90 27.265v45.294L66.272 61.337z"/>`),
		g.Group(children),
	)
}