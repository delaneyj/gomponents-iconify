package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 8a9 9 0 0 0-9 9v1H1v-1C1 10.925 5.925 6 12 6s11 4.925 11 11v1h-2v-1a9 9 0 0 0-9-9Zm0 3a6 6 0 0 0-6 6v1H4v-1a8 8 0 1 1 16 0v1h-2v-1a6 6 0 0 0-6-6Zm0 3a3 3 0 0 0-3 3v1H7v-1a5 5 0 0 1 10 0v1h-2v-1a3 3 0 0 0-3-3Z"/>`),
		g.Group(children),
	)
}