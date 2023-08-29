package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M30 101h708c16 0 30 14 30 30v432c0 16-14 30-30 30H30c-16 0-30-14-30-30V131c0-16 14-30 30-30zm21 440h666V152H51v389zm72-348h61v62h-61v-62zm116 0h61v62h-61v-62zm114 0h62v62h-62v-62zm116 0h60v62h-60v-62zm115 0h61v62h-61v-62zM181 316h62v62h-62v-62zm116 62v-62h61v62h-61zm115 0v-62h61v62h-61zm177 0h-62v-62h62v62zm-60 122H239v-61h290v61z"/>`),
		g.Group(children),
	)
}