package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Raindrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.25 19.855h0a9.75 9.75 0 0 1 9.75 9.75v9.75h0h-9.75a9.75 9.75 0 0 1-9.75-9.75v0a9.75 9.75 0 0 1 9.75-9.75Zm9.75 19.5h9.75a9.75 9.75 0 0 0 9.75-9.75h0a9.75 9.75 0 0 0-9.75-9.75h0a9.75 9.75 0 0 0-9.75 9.75"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.995 12.371h0a12.72 12.72 0 0 1 0 17.989L24 39.354h0l-8.994-8.994a12.72 12.72 0 0 1 0-17.989h0a12.72 12.72 0 0 1 17.989 0Z"/>`),
		g.Group(children),
	)
}