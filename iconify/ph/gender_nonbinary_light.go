package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderNonbinaryLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M134 98.27V58.6l30.91 18.54a6 6 0 1 0 6.18-10.28L139.66 48l31.43-18.85a6 6 0 1 0-6.18-10.29L128 41L91.09 18.86a6 6 0 1 0-6.18 10.29L116.34 48L84.91 66.86a6 6 0 1 0 6.18 10.28L122 58.6v39.67a70 70 0 1 0 12 0ZM128 226a58 58 0 1 1 58-58a58.07 58.07 0 0 1-58 58Z"/>`),
		g.Group(children),
	)
}