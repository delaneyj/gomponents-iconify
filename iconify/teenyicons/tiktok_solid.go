package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TiktokSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 0h1v1a4 4 0 0 0 4 4v1a4.992 4.992 0 0 1-4-2v7a4 4 0 1 1-4-4v1a3 3 0 1 0 3 3V0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}