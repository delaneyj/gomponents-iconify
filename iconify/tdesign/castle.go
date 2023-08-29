package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 2v2h2V2h2v2h2V2h2v2h2V2h2v12h2v-1h2v9H1v-9h2v1h2V2h2ZM5 16H3v4h2v-4Zm2 4h2v-7h6v7h2V6H7v14Zm12 0h2v-4h-2v4Zm-6 0v-5h-2v5h2Z"/>`),
		g.Group(children),
	)
}