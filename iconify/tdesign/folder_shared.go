package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderShared(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1V2.5Zm2 2V19h18V7H10.52l-2-2.5H3Zm9 6a1 1 0 1 1 0 2a1 1 0 0 1 0-2Zm2.4 2.8a3 3 0 1 0-4.8 0A3.994 3.994 0 0 0 8 16.5v1h2v-1a2 2 0 1 1 4 0v1h2v-1c0-1.309-.628-2.47-1.6-3.2Z"/>`),
		g.Group(children),
	)
}