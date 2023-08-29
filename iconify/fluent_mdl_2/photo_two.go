package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1536H0V256h2048zM128 384v677l448-447l640 640l256-256l448 447V384H128zm0 1280h1317L576 794l-448 449v421zm1792 0v-37l-448-449l-166 166l321 320h293zm-320-896q-26 0-45-19t-19-45q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19z"/>`),
		g.Group(children),
	)
}