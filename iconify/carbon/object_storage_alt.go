package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ObjectStorageAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 20h-2v2h2v6H4v-6h10v-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2Z"/><circle cx="7" cy="25" r="1" fill="currentColor"/><path fill="currentColor" d="M21 14a2.981 2.981 0 0 0-2.037.811l-4.004-2.402A2.958 2.958 0 0 0 15 12a2.958 2.958 0 0 0-.041-.409l4.004-2.402A2.99 2.99 0 1 0 18 7a2.934 2.934 0 0 0 .042.41l-4.004 2.402a3 3 0 1 0 0 4.377l4.003 2.402A2.934 2.934 0 0 0 18 17a3 3 0 1 0 3-3Zm0-8a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm-9 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm9 5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}