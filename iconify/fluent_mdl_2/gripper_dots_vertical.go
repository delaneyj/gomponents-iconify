package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GripperDotsVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1152 384h256v256h-256V384zm0 768V896h256v256h-256zm0 512v-256h256v256h-256zM640 640V384h256v256H640zm0 512V896h256v256H640zm0 512v-256h256v256H640z"/>`),
		g.Group(children),
	)
}