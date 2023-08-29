package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MacCommand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 13a4 4 0 0 0 4-4V8a4 4 0 0 0-4-4h-1a4 4 0 0 0-4 4v3h-6V8a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v1a4 4 0 0 0 4 4h3v6H8a4 4 0 0 0-4 4v1a4 4 0 0 0 4 4h1a4 4 0 0 0 4-4v-3h6v3a4 4 0 0 0 4 4h1a4 4 0 0 0 4-4v-1a4 4 0 0 0-4-4h-3v-6Zm-3-5a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-3ZM8 11a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v3H8Zm3 13a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2h3Zm8-5h-6v-6h6Zm2 2h3a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2Z"/>`),
		g.Group(children),
	)
}