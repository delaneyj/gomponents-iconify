package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1920h384v128H512v-128h384V896H231L530 0h860l299 896h-409v192q0 26-19 45t-45 19q-26 0-45-19t-19-45V896h-128v1024zM409 768h1102l-213-640H622L409 768z"/>`),
		g.Group(children),
	)
}