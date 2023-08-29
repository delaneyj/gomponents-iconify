package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Catalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 2H8a2 2 0 0 0-2 2v4H4v2h2v5H4v2h2v5H4v2h2v4a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Zm0 26H8v-4h2v-2H8v-5h2v-2H8v-5h2V8H8V4h18Z"/><path fill="currentColor" d="M14 8h8v2h-8zm0 7h8v2h-8zm0 7h8v2h-8z"/>`),
		g.Group(children),
	)
}