package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chair(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 13V4c0-1.103-.897-2-2-2H7c-1.103 0-2 .897-2 2v9a1 1 0 0 0-1 1v8h2v-5h12v5h2v-8a1 1 0 0 0-1-1zm-2-9v9h-2V4h2zm-4 0v9h-2V4h2zM7 4h2v9H7V4z"/>`),
		g.Group(children),
	)
}