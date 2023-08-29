package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 128h1664v1664H128V128zm1536 128H256v640h1408V256zM256 1024v640h640v-640H256zm768 640h640v-640h-640v640z"/>`),
		g.Group(children),
	)
}