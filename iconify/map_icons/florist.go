package map_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Florist(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 50 50"),
		g.Raw(`<ellipse cx="25.015" cy="25.003" fill="none" rx="2.674" ry="2.688"/><path fill="currentColor" d="M25.554 36.172c3.601 16.811 23.919 3.049 3.74-8.661c22.24 12.908 22.293-17.916.045-5.018c22.248-12.897-4.317-28.315-4.322-2.509c.005-25.806-26.564-10.4-4.327 2.509c-22.237-12.91-22.263 17.915-.019 5.016c-20.19 11.704.082 25.482 3.684 8.637l1.199.026zm-3.213-11.169a2.68 2.68 0 0 1 2.674-2.688a2.68 2.68 0 0 1 2.673 2.688a2.68 2.68 0 0 1-2.673 2.688a2.683 2.683 0 0 1-2.674-2.688z"/>`),
		g.Group(children),
	)
}