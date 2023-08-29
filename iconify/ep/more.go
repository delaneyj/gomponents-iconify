package ep

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func More(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M176 416a112 112 0 1 0 0 224a112 112 0 0 0 0-224m0 64a48 48 0 1 1 0 96a48 48 0 0 1 0-96zm336-64a112 112 0 1 1 0 224a112 112 0 0 1 0-224zm0 64a48 48 0 1 0 0 96a48 48 0 0 0 0-96zm336-64a112 112 0 1 1 0 224a112 112 0 0 1 0-224zm0 64a48 48 0 1 0 0 96a48 48 0 0 0 0-96z"/>`),
		g.Group(children),
	)
}