package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContextMenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 128v1664H256V128h1536zm-128 128H384v1408h1280V256zm-256 768H640V896h768v128zm0 384H640v-128h768v128zm0-768H640V512h768v128z"/>`),
		g.Group(children),
	)
}