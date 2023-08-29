package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EdgeNode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 6h4v4h-4z"/><circle cx="7" cy="7" r="1" fill="currentColor"/><circle cx="25" cy="25" r="1" fill="currentColor"/><circle cx="25" cy="21" r="1" fill="currentColor"/><circle cx="25" cy="17" r="1" fill="currentColor"/><path fill="currentColor" d="M22 17v-2h-2v-1a2.002 2.002 0 0 0-2-2h-1v-2h-2v2h-2v-2h-2v2h-1a2.002 2.002 0 0 0-2 2v1H6v2h2v2H6v2h2v1a2.002 2.002 0 0 0 2 2h1v2h2v-2h2v2h2v-2h1a2.002 2.002 0 0 0 2-2v-1h2v-2h-2v-2Zm-4 5h-8v-8h8Z"/><path fill="currentColor" d="M28 30H4a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v24a2.002 2.002 0 0 1-2 2ZM4 4v24h24V4Z"/>`),
		g.Group(children),
	)
}