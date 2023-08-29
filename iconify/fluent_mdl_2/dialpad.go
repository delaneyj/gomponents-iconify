package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dialpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 384H384V128h256v256zm512 0H896V128h256v256zm512 0h-256V128h256v256zM640 896H384V640h256v256zm512 0H896V640h256v256zm512 0h-256V640h256v256zM640 1408H384v-256h256v256zm512 0H896v-256h256v256zm0 512H896v-256h256v256zm512-512h-256v-256h256v256z"/>`),
		g.Group(children),
	)
}