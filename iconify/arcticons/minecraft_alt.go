package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MinecraftAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" class="c"><path d="m40.9 14.273l-.055 19.454L24.027 43.5l.056-19.453L40.9 14.273zm-16.817 9.774L24.027 43.5L7.1 33.727l.055-19.454l16.928 9.774z"/><path d="m40.9 14.273l-16.817 9.774l-16.928-9.774L23.973 4.5L40.9 14.273z"/></g>`),
		g.Group(children),
	)
}