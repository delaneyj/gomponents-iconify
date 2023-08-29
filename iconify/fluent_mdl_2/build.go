package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Build(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 1280H384v-256h256v256zm1280-512v1152H128V768h128v640h1536V768h128zM896 1536H640v256h256v-256zm512 0h-256v256h256v-256zm-512-512h256v256H896v-256zm768 256h-256v-256h256v256zM960 922L659 621l90-90l147 146V0h128v677l147-146l90 90l-301 301z"/>`),
		g.Group(children),
	)
}