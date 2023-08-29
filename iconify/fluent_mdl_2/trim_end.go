package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TrimEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M384 0h1408v2048H384v-128h768V128H384V0zm1280 1920V128h-384v1792h384z"/>`),
		g.Group(children),
	)
}