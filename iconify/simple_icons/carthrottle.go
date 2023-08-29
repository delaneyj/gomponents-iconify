package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Carthrottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 19.99h5.31l1-5.76h2.673L7.97 19.99h5.272l1.037-5.76h2.824l-1 5.76h7.584L21.9 17.029L24 4.01h-5.16l-.987 5.647h-2.86l.936-5.647H8.483l1.724 2.749l-.487 2.898H6.996l.9-5.647H.35l1.76 2.774Z"/>`),
		g.Group(children),
	)
}