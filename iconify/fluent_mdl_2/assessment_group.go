package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AssessmentGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 1536H384v-256h256v256zm512 0H896v-256h256v256zm512-1024h-256V256h256v256zm0 1024h-256v-256h256v256zm-512-512H896V768h256v256zm512 0h-256V768h256v256zm256-768v1536H128V256h128v1408h1536V256h128z"/>`),
		g.Group(children),
	)
}