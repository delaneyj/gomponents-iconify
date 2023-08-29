package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TagSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M10 4.5a.5.5 0 1 1 1 0a.5.5 0 0 1-1 0Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.146.146A.5.5 0 0 1 7.5 0h6A1.5 1.5 0 0 1 15 1.5v6a.5.5 0 0 1-.146.354l-7 7a.5.5 0 0 1-.708 0l-7-7a.5.5 0 0 1 0-.708l7-7ZM10.5 3a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}