package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 0v4h5.5A1.5 1.5 0 0 1 15 5.5v7a1.5 1.5 0 0 1-1.5 1.5h-12A1.5 1.5 0 0 1 0 12.5v-7A1.5 1.5 0 0 1 1.5 4H7V0h1Zm2 6h1v1h-1V6Zm1 2h-1v1h1V8Zm0 3H4v1h7v-1ZM7 8h1v1H7V8ZM5 8H4v1h1V8Zm3-2H7v1h1V6ZM4 6h1v1H4V6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}