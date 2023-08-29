package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2011 960l-446 445l-90-90l291-291H768V896h998l-291-291l90-90l446 445zm-859 320h128v640H0V0h1280v640h-128V128H128v1664h1024v-512z"/>`),
		g.Group(children),
	)
}