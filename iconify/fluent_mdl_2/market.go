package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Market(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 384v640h-128V603l-768 768l-384-384l-675 674l-90-90l765-766l384 384l677-677h-421V384h640z"/>`),
		g.Group(children),
	)
}