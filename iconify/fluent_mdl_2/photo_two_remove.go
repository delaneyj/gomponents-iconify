package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoTwoRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1408 576q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45zm410 1024l227 227l-90 90l-227-227l-227 227l-90-90l227-227l-227-227l90-90l227 227l227-227l90 90l-227 227zm-361 0l-65 64H0V256h1792v1008l-64 65l-64-65V384H128v421l192-191l512 512l256-256l323 322l-91 91l-232-233l-166 166l321 320h149l65 64zm-396-64L320 794L128 987v549h933z"/>`),
		g.Group(children),
	)
}