package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColumnLeftTwoThirds(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1792V256h1280v1536H0zM128 384v1280h1024V384H128zm1280-128h640v1536h-640V256zm512 1408V384h-384v1280h384z"/>`),
		g.Group(children),
	)
}