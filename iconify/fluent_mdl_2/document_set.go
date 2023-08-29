package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentSet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1280h128v768H128V640h128V256h128V0h1408v1280zM512 128v128h1024v1024h128V128H512zM384 384v256h347l640 640h37V384H384zm1408 1536v-512h-475L677 768H256v1152h1536z"/>`),
		g.Group(children),
	)
}