package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Column(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 0h1024v1920H512V0zm896 1792V384H640v1408h768zm0-1536V128H640v128h768z"/>`),
		g.Group(children),
	)
}