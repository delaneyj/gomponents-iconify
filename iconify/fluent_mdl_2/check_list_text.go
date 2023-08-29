package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckListText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 1664v-128h1408v128H640zm0-384v-128h1408v128H640zm0-896h1408v128H640V384zm0 512V768h1408v128H640z"/>`),
		g.Group(children),
	)
}