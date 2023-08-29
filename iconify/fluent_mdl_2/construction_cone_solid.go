package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConstructionConeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 1920v128H128v-128h270l96-384h802l-32-128H526l64-256h610l-32-128H622l224-896h356l448 1792h270z"/>`),
		g.Group(children),
	)
}