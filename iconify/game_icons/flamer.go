package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flamer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M166 121c0 90 90 105 90 180c0 30-30 75-75 75s-75-45-45-120c-45 30-60 60-60 90c0 75 75 150 180 150s180-45 180-135c.67-133.125-153.4-177.596-195-240c-30-45-15-75 15-105c-60 15-90 57-90 105z"/>`),
		g.Group(children),
	)
}