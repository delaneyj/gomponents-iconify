package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriggerHurt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M254.25 121.03A60 60 0 0 0 196 181a60 60 0 0 0 120 0a60 60 0 0 0-61.75-59.97zM136 271l-60 60H46c-15 0-30 15-30 30v30h480v-30c0-15-15-30-30-30h-30l-60-60l-60 60l-60-60l-60 60l-60-60z"/>`),
		g.Group(children),
	)
}