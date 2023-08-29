package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageAcount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h2048v2048H0V0zm1920 1920V512H128v1408h1792zm0-1536V128H128v256h1792zm-128 256v512H256V640h1536zm-128 384V768H384v256h1280zm128 256v512H256v-512h1536zm-128 384v-256H384v256h1280z"/>`),
		g.Group(children),
	)
}