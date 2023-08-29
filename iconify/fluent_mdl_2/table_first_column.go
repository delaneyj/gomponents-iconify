package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableFirstColumn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 128h2048v1664H0V128zm1920 512V256h-512v384h512zM768 256v384h512V256H768zm0 1024v384h512v-384H768zm0-128h512V768H768v384zm640-384v384h512V768h-512zm0 896h512v-384h-512v384z"/>`),
		g.Group(children),
	)
}