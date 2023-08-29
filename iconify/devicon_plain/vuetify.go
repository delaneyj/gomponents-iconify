package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vuetify(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M2.668 10.668L33.34 63.797l1.605 2.793l22.883-39.65l9.399-16.273Zm76.336 0L65.297 34.414L40.84 76.79L64 116.922l30.672-53.125l30.66-53.129Zm0 0"/>`),
		g.Group(children),
	)
}