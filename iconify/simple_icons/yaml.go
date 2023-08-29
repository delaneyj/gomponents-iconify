package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yaml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 .97l4.111 6.453v4.09h2.638v-4.09L11.053.969H8.214L5.58 5.125L2.965.969Zm12.093.024l-4.47 10.544h2.114l.97-2.345h4.775l.804 2.345h2.26L14.255.994Zm1.133 2.225l1.463 3.87h-3.096zm3.06 9.475v10.29H24v-2.199h-5.454v-8.091zm-12.175.002v10.335h2.217v-7.129l2.32 4.792h1.746l2.4-4.96v7.295h2.127V12.696h-2.904L9.44 17.37l-2.455-4.674Z"/>`),
		g.Group(children),
	)
}