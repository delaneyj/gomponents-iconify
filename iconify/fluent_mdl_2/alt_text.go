package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AltText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1536 704q0-26 19-45t45-19q26 0 45 19t19 45q0 26-19 45t-45 19q-26 0-45-19t-19-45zM0 256h2048v1024h-128V384H128v677l448-447l640 640l256-256l283 282h-182l-101-102l-101 102h-310L576 794l-448 449v421h512v128H0V256zm768 1792v-640h1280v640H768zm128-512v384h1024v-384H896zm256 256v-128h512v128h-512z"/>`),
		g.Group(children),
	)
}