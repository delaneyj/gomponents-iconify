package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewAll(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 128h1664v1664H128V128zm1536 768V256h-640v640h640zM896 256H256v640h640V256zm-640 768v640h640v-640H256zm768 640h640v-640h-640v640z"/>`),
		g.Group(children),
	)
}