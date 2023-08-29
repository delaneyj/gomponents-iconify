package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Migrate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 2H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h9v6.17l-2.59-2.58L11 15l5 5l5-5l-1.41-1.41L17 16.17V10h9a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zM6 4h4v4H6zm20 4H12V4h14zm0 14H6a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2zM6 24h14v4H6zm20 4h-4v-4h4z"/>`),
		g.Group(children),
	)
}