package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m448 102l365 365l-90 90l-211-210v549H384V347L173 557l-90-90l365-365zm275 1389l90 90l-365 365l-365-365l90-90l211 210v-549h128v549l211-210zM2048 256v1536H896v-128h1024V384H896V256h1152zm-384 256l-197 320h340l-626 704H927l165-384H856l320-640h488zm-425 448l196-320h-179l-192 384h222l-163 384l398-448h-282z"/>`),
		g.Group(children),
	)
}