package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExternalXAML(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 256H256v1664H128V128h1792v1024h-128V256zm-137 1408l-220 384H997l-220-384l220-384h438l220 384zm361 0l-219 384h-148l220-384l-220-384h148l219 384zm-1453 0l220 384H635l-219-384l219-384h148l-220 384z"/>`),
		g.Group(children),
	)
}