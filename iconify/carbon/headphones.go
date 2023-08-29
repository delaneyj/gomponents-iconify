package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M25 16v-1a9 9 0 0 0-18 0v1a5 5 0 0 0 0 10h2V15a7 7 0 0 1 14 0v11h2a5 5 0 0 0 0-10ZM4 21a3 3 0 0 1 3-3v6a3 3 0 0 1-3-3Zm21 3v-6a3 3 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}