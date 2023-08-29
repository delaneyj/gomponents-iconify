package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnRightTwoThirds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M768 256h1280v1536H768V256zm1152 1408V384H896v1280h1024zM0 1792V256h640v1536H0zM128 384v1280h384V384H128z"/>`),
		g.Group(children),
	)
}