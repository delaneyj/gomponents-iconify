package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M8 11h3v12h2V11h3V9H8v2z" fill="currentColor"/><path d="M24 9h-2V7h-2v2h-2v2h2v6l1 1l1-1v-6h2V9z" fill="currentColor"/>`),
		g.Group(children),
	)
}