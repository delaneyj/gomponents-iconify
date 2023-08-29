package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataAccessor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 20h-2v2h2v6H4v-6h2v-2H4a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h24a2.002 2.002 0 0 0 2-2v-6a2.002 2.002 0 0 0-2-2Z"/><circle cx="7" cy="25" r="1" fill="currentColor"/><path fill="currentColor" d="M21 13a2.96 2.96 0 0 0-1.285.3l-2.3-2.3l2.3-2.3A2.96 2.96 0 0 0 21 9a3 3 0 1 0-3-3a2.966 2.966 0 0 0 .3 1.285l-2.3 2.3l-2.3-2.3A2.966 2.966 0 0 0 14 6a3 3 0 1 0-3 3a2.96 2.96 0 0 0 1.285-.3l2.3 2.3l-2.3 2.3A2.96 2.96 0 0 0 11 13a3 3 0 1 0 3 3a2.966 2.966 0 0 0-.3-1.285l2.3-2.3l2.3 2.3A2.966 2.966 0 0 0 18 16a3 3 0 1 0 3-3Zm0-8a1 1 0 1 1-1 1a1 1 0 0 1 1-1ZM10 6a1 1 0 1 1 1 1a1 1 0 0 1-1-1Zm1 11a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm10 0a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}