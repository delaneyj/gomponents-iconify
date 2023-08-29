package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RowsChild(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 1024v640H768v-256H256V896H0V256h1280v640H384v384h384v-256h1280zM128 768h1024V384H128v384zm1792 384H896v384h1024v-384z"/>`),
		g.Group(children),
	)
}