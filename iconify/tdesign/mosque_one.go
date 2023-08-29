package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MosqueOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 1.586l4 4v2.31c.124-.21.26-.418.412-.62c.785-1.05 1.957-1.938 3.588-2.198V3.5h2v1.578c1.63.26 2.802 1.148 3.588 2.198c.658.88 1.05 1.873 1.252 2.724H22v12H2V5.586l4-4ZM18.764 10a5.346 5.346 0 0 0-.777-1.526C17.367 7.644 16.437 7 15 7c-1.438 0-2.366.645-2.987 1.474A5.346 5.346 0 0 0 11.236 10h7.528ZM10 12v8h2v-6h6v6h2v-8H10Zm6 8v-4h-2v4h2Zm-8 0V6.414l-2-2l-2 2V20h4ZM7.004 8v2.004H5v-2h.004V8h2Z"/>`),
		g.Group(children),
	)
}