package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImagePixel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1536H0V256h2048zm-128 128H128v512h256v256H128v512h1024v-256h256v256h256v-256h256V384zM640 896H384V640h256v256zm0 256V896h256v256H640zm512 0v256H896v-256h256zm512 0v256h-256v-256h256zm0-256h-256V640h256v256z"/>`),
		g.Group(children),
	)
}