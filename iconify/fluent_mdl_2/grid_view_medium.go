package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridViewMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 128v1664H128V128h1664zm-128 128h-640v640h640V256zm-1408 0v640h640V256H256zm0 1408h640v-640H256v640zm1408 0v-640h-640v640h640z"/>`),
		g.Group(children),
	)
}