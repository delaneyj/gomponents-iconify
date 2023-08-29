package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalDistributeCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 256h256v1536H640v256H512v-256H256V256h256V0h128v256zm128 128H384v1280h384V384zm1024 1152h-256v512h-128v-512h-256V512h256V0h128v512h256v1024zm-128-896h-384v768h384V640z"/>`),
		g.Group(children),
	)
}