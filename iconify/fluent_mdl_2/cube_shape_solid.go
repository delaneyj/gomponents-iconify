package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CubeShapeSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m960 152l750 376l-750 375l-750-375l750-376zM128 1607V665l768 384v942l-768-384zm896-558l768-384v942l-768 384v-942z"/>`),
		g.Group(children),
	)
}