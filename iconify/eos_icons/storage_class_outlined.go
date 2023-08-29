package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StorageClassOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m23 18l-2.4 3.2a1.816 1.816 0 0 1-1.6.8h-7a1.003 1.003 0 0 1-1-1v-6a1.003 1.003 0 0 1 1-1h7a2.088 2.088 0 0 1 1.6.8ZM21 6H3V2h18ZM2 1v6h20V1"/><path fill="currentColor" d="M4 3h2v2H4zM3 14v-4h18v3.33l1 1.337V9H2v6h7v-1H3z"/><path fill="currentColor" d="M4 11h2v2H4zm0 8h2v2H4z"/><path fill="currentColor" d="M3 22v-4h6v-1H2v6h7v-1H3z"/>`),
		g.Group(children),
	)
}