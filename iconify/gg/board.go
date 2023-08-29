package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Board(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 4a4 4 0 0 0-4 4v8a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4V8a4 4 0 0 0-4-4H6Zm8 2h-4v12h4V6Zm2 0v12h2a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-2ZM6 18h2V6H6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}