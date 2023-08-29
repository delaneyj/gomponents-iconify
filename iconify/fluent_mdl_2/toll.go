package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 1920h128v128H0V0h1024v384H768v920l1085-542l185 371l-1270 635v152zm0-296l256-128v-176l-256 128v176zm384-192l256-128v-176l-256 128v176zm643-498l-259 130v176l330-165l-71-141zM640 1280V640H384v640h256zM128 256h768V128H128v128zm512 1664v-512H256V512h384V384H128v1536h512z"/>`),
		g.Group(children),
	)
}