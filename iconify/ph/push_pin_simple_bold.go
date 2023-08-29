package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PushPinSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 164h-5.93L190.3 52h1.7a12 12 0 0 0 0-24H64a12 12 0 0 0 0 24h1.7L45.93 164H40a12 12 0 0 0 0 24h76v52a12 12 0 0 0 24 0v-52h76a12 12 0 0 0 0-24ZM90.07 52h75.86l19.77 112H70.3Z"/>`),
		g.Group(children),
	)
}