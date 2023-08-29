package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMusic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 3a2 2 0 0 1 2 2v16l-7-3l-7 3V5a2 2 0 0 1 2-2h10m-6 8a2 2 0 0 0-2 2a2 2 0 0 0 2 2a2 2 0 0 0 2-2V8h3V6h-4v5.27c-.29-.17-.64-.27-1-.27Z"/>`),
		g.Group(children),
	)
}