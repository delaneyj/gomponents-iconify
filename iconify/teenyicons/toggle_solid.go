package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToggleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 3.5A3.5 3.5 0 0 1 3.5 0h8a3.5 3.5 0 1 1 0 7h-8A3.5 3.5 0 0 1 0 3.5ZM3.5 2a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3ZM15 11.5a3.5 3.5 0 0 1-3.5 3.5h-8a3.5 3.5 0 1 1 0-7h8a3.5 3.5 0 0 1 3.5 3.5ZM11.5 13a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}