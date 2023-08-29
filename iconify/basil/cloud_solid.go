package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.932 4.708c-1.912 0-3.59.992-4.552 2.486A5.875 5.875 0 1 0 6.875 18.75h11.648a4.478 4.478 0 1 0-.19-8.951a5.41 5.41 0 0 0-5.401-5.09Z"/>`),
		g.Group(children),
	)
}