package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderInnerSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 2H1V1h1v1Zm3 0H4V1h1v1Zm2 5H1v1h6v6h1V8h6V7H8V1H7v6Zm4-5h-1V1h1v1Zm3 0h-1V1h1v1ZM2 5H1V4h1v1Zm12 0h-1V4h1v1ZM2 11H1v-1h1v1Zm12 0h-1v-1h1v1ZM2 14H1v-1h1v1Zm3 0H4v-1h1v1Zm6 0h-1v-1h1v1Zm3 0h-1v-1h1v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}