package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Intersect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22 24h-4v-2h4v-4h2v4a2.002 2.002 0 0 1-2 2zM10 14H8v-4a2.002 2.002 0 0 1 2-2h4v2h-4z"/><path fill="currentColor" d="M28 8h-4V4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v18a2.002 2.002 0 0 0 2 2h4v4a2.002 2.002 0 0 0 2 2h18a2.002 2.002 0 0 0 2-2V10a2.002 2.002 0 0 0-2-2Zm0 20H10v-4h4v-2h-4v-4H8v4H4V4h18v4h-4v2h4v4h2v-4h4Z"/>`),
		g.Group(children),
	)
}