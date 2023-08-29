package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 6h-4V2h-2v4h-4v2h4v4h2V8h4V6zm-6 9v7.556A3.955 3.955 0 0 0 22 22a4 4 0 1 0 4 4V15zm-2 13a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zM17 6h-7a2.002 2.002 0 0 0-2 2v14.556A3.956 3.956 0 0 0 6 22a4 4 0 1 0 4 4V8h7zM6 28a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}