package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imgur(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 4.5a5.58 5.58 0 0 1 5.58 5.58h0A5.58 5.58 0 0 1 24 15.66h0a5.58 5.58 0 0 1-5.58-5.58h0A5.58 5.58 0 0 1 24 4.5Zm0 15.13a5.79 5.79 0 0 1 5.58 6v11.86a5.79 5.79 0 0 1-5.58 6a5.79 5.79 0 0 1-5.58-6V25.64A5.79 5.79 0 0 1 24 19.63Z"/>`),
		g.Group(children),
	)
}