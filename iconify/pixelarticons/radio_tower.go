package pixelarticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioTower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2h-2v2h2v12h-2v2h2v-2h2V4h-2V2ZM2 4H0v12h2v2h2v-2H2V4Zm0 0V2h2v2H2Zm4 2H4v8h2V6Zm0 0V4h2v2H6Zm4 0h4v2h-4V6Zm0 6H8V8h2v4Zm4 0h-4v2H8v4H6v4h2v-4h2v-4h4v4h2v4h2v-4h-2v-4h-2v-2Zm0 0h2V8h-2v4Zm6-6h-2V4h-2v2h2v8h2V6Z"/>`),
		g.Group(children),
	)
}