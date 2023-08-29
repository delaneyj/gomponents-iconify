package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Promotion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m64 448l832-320l-128 704l-446.08-243.328L832 192L242.816 545.472L64 448zm256 512V657.024L512 768L320 960z"/>`),
		g.Group(children),
	)
}