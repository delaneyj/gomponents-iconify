package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Crossroad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m164 16l-32 16l32 32l64 16V32l-64-16zm184 0l-64 16v48l64-16l32-32l-32-16zM247 32v112h18V32h-18zM16 64l182.7 182.7L96 496h320L313.3 246.7L496 64h-64L256 192L80 64H16z"/>`),
		g.Group(children),
	)
}