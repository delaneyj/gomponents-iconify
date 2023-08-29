package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntryDecline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M512 768h1024v128H512V768zm1024-256H512V384h1024v128zm-384 1408l127 128H256V0h1536v1348l-64 63l-64-64V128H384v1792h768zm576 125l3 3h-6l3-3zm-192-893l-129 128H512v-128h1024zm-317 384l128 128H512v-128h707zm600 192l226 227l-90 90l-227-226l-227 227l-90-91l227-227l-227-227l90-90l227 227l227-227l90 91l-226 226z"/>`),
		g.Group(children),
	)
}