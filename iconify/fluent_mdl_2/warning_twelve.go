package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1884 2048H-7L939 57l945 1991zM263 1877h1351L939 455L263 1877zm761-1024v555H853V853h171zm-171 683h171v171H853v-171z"/>`),
		g.Group(children),
	)
}