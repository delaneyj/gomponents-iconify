package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M32 30.5v-29c0-.827-.673-1.5-1.5-1.5h-29C.673 0 0 .673 0 1.5v9a.5.5 0 0 0 1 0v-9a.5.5 0 0 1 .5-.5h29a.5.5 0 0 1 .5.5v29a.5.5 0 0 1-.5.5h-9a.5.5 0 0 0 0 1h9c.827 0 1.5-.673 1.5-1.5z"/><path d="M1.5 13c-.827 0-1.5.673-1.5 1.5v16c0 .827.673 1.5 1.5 1.5h16c.827 0 1.5-.673 1.5-1.5V13.707l7-7V13.5a.5.5 0 0 0 1 0v-8a.5.5 0 0 0-.5-.5h-8a.5.5 0 0 0 0 1h6.792l-7 7H1.5zM18 30.5a.5.5 0 0 1-.5.5h-16a.5.5 0 0 1-.5-.5v-16a.5.5 0 0 1 .5-.5H18v16.5z"/></g>`),
		g.Group(children),
	)
}