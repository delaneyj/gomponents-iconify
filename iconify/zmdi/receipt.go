package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Receipt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M320 323v-43H64v43h256zm0-86v-42H64v42h256zm0-85v-43H64v43h256zM0 429V3l32 32L64 3l32 32l32-32l32 32l32-32l32 32l32-32l32 32l32-32l32 32l32-32v426l-32-32l-32 32l-32-32l-32 32l-32-32l-32 32l-32-32l-32 32l-32-32l-32 32l-32-32z"/>`),
		g.Group(children),
	)
}