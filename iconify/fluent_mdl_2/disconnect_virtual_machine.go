package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DisconnectVirtualMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1277 1501l-227 227l227 227l-90 90l-227-227l-227 227l-90-90l227-227l-227-227l90-90l227 227l227-227l90 90zm-669 163l64 64l-64 64H128v-128h480zm640 64l64-64h480v128h-480l-64-64zM128 128h1664v1152h-768v96l-64 64l-64-64v-96H128V128zm1536 1024V256H256v896h1408z"/>`),
		g.Group(children),
	)
}