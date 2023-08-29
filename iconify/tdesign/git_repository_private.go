package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitRepositoryPrivate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 7.5a6.5 6.5 0 0 1 13 0V9H21v14H3V9h2.5V7.5Zm2 1.5h9V7.5a4.5 4.5 0 1 0-9 0V9ZM5 11v10h14V11H5Zm2 1.504h2.004V15H7v-2.496ZM7 16.5h2.004v2.504H7V16.5Z"/>`),
		g.Group(children),
	)
}