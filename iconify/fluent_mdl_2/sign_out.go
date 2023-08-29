package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 1792h896v128H128V0h1024v128H256v1664zm1679-832l-487 488l-80-80l345-344H640V896h1073l-345-344l80-80l487 488z"/>`),
		g.Group(children),
	)
}