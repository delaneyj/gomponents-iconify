package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 1792V256h896v1536H0zM128 384v1280h640V384H128zm896-128h896v1536h-896V256zm768 1408V384h-640v1280h640z"/>`),
		g.Group(children),
	)
}