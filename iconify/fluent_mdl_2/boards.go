package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1664 512h256v1536H640v-256H384v-256H128V0h1280v256h256v256zM384 1408V256h896V128H256v1280h128zm256 256V512h896V384H512v1280h128zm1152 256V640H768v1280h1024z"/>`),
		g.Group(children),
	)
}