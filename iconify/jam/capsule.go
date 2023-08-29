package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Capsule(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9V5a3 3 0 1 0-6 0v4h6zm0 2H2v4a3 3 0 0 0 6 0v-4zM5 0a5 5 0 0 1 5 5v10a5 5 0 0 1-10 0V5a5 5 0 0 1 5-5z"/>`),
		g.Group(children),
	)
}