package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlobStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h2048v2048H0V0zm1920 1920V512H128v1408h1792zm0-1536V128H128v256h1792zM768 1280v512H256v-512h512zm-128 384v-256H384v256h256zM1792 640v512h-512V640h512zm-128 384V768h-256v256h256zm-768 768v-512h896v512H896zm128-384v256h640v-256h-640zm128-768v512H256V640h896zm-128 384V768H384v256h640z"/>`),
		g.Group(children),
	)
}