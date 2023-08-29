package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitCoinFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.005 22.003c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Zm-1-6v2h2v-2h1a2.5 2.5 0 0 0 2-4a2.5 2.5 0 0 0-2-4h-1v-2h-2v2h-3v8h3Zm-1-3h4a.5.5 0 0 1 0 1h-4v-1Zm0-3h4a.5.5 0 0 1 0 1h-4v-1Z"/>`),
		g.Group(children),
	)
}