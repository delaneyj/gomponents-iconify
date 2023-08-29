package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderLeftSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 14V1h1v13H1ZM5 2H4V1h1v1Zm3 0H7V1h1v1Zm3 0h-1V1h1v1Zm3 0h-1V1h1v1ZM8 5H7V4h1v1Zm6 0h-1V4h1v1ZM5 8H4V7h1v1Zm3 0H7V7h1v1Zm3 0h-1V7h1v1Zm3 0h-1V7h1v1Zm-6 3H7v-1h1v1Zm6 0h-1v-1h1v1Zm-9 3H4v-1h1v1Zm3 0H7v-1h1v1Zm3 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}