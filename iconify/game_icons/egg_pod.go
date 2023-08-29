package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EggPod(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16c-67.5 0-105 120-105 210c0 45 45 90 105 90s105-45 105-90c0-90-37.5-210-105-210zM136 271c0 90 45 75 45 135c0 90-75 60-75 90h300c0-30-75 0-75-90c0-60 45-45 45-135c-90 105-150 105-240 0z"/>`),
		g.Group(children),
	)
}