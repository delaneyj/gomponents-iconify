package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RssThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M96.08 159.92A67.52 67.52 0 0 1 116 208a4 4 0 0 1-8 0a60 60 0 0 0-60-60a4 4 0 0 1 0-8a67.52 67.52 0 0 1 48.08 19.92ZM48 92a4 4 0 0 0 0 8a108 108 0 0 1 108 108a4 4 0 0 0 8 0A116 116 0 0 0 48 92Zm116 0A162.92 162.92 0 0 0 48 44a4 4 0 0 0 0 8a155 155 0 0 1 110.31 45.69A155 155 0 0 1 204 208a4 4 0 0 0 8 0a162.92 162.92 0 0 0-48-116ZM52 196a8 8 0 1 0 8 8a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}