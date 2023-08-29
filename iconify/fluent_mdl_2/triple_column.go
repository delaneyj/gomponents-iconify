package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TripleColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1792V256h512v1536H128zM256 384v1280h256V384H256zm512 1408V256h512v1536H768zM896 384v1280h256V384H896zm512-128h512v1536h-512V256zm384 1408V384h-256v1280h256z"/>`),
		g.Group(children),
	)
}