package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleTarget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M247 65.16L32.34 440.8l61.79-35.7L247 137.6zm18 .38V137l158.3 271.3l62.7 36.1C412.2 318.2 338.6 191.8 265 65.54zM415.4 424.5l-321.3 1.4l-62.72 36.2l445.82-1.9z"/>`),
		g.Group(children),
	)
}