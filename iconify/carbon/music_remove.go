package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26.41 7L30 3.41L28.59 2L25 5.59L21.41 2L20 3.41L23.59 7L20 10.59L21.41 12L25 8.41L28.59 12L30 10.59L26.41 7zM24 15v7.556A3.955 3.955 0 0 0 22 22a4 4 0 1 0 4 4V15zm-2 13a2 2 0 1 1 2-2a2.003 2.003 0 0 1-2 2zM17 6h-7a2.002 2.002 0 0 0-2 2v14.556A3.956 3.956 0 0 0 6 22a4 4 0 1 0 4 4V8h7zM6 28a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2z"/>`),
		g.Group(children),
	)
}