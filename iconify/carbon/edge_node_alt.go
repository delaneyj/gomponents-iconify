package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeNodeAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 22a2.981 2.981 0 0 0-2.037.811l-4.004-2.402a2.043 2.043 0 0 0 0-.818l4.004-2.402A2.99 2.99 0 1 0 24 15a2.934 2.934 0 0 0 .042.41l-4.004 2.401a3 3 0 1 0 0 4.377l4.003 2.403A2.934 2.934 0 0 0 24 25a3 3 0 1 0 3-3Zm0-8a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm-9 7a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm9 5a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/><circle cx="7" cy="8" r="1" fill="currentColor"/><circle cx="7" cy="16" r="1" fill="currentColor"/><circle cx="7" cy="24" r="1" fill="currentColor"/><path fill="currentColor" d="M22 13V5a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v22a2.002 2.002 0 0 0 2 2h16v-2H4v-6h8v-2H4v-6ZM4 5h16v6H4Z"/>`),
		g.Group(children),
	)
}