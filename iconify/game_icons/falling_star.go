package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FallingStar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m166.4 25.64l-12.8 12.72l160 160.04l12.8-12.8l-160-159.96zm-32 95.96L128 128l-6.4 6.4l160 160l12.8-12.8l-160-160zm-96.08 32L25.6 166.4l160 160l12.8-12.8l-160.08-160zm314.78 86.6l-29.4 84.1l-85.4 26l71 54l-1.7 89.2l73.2-50.8l84.4 29.1l-25.7-85.3l53.8-71.2l-89.1-2l-51.1-73.1z"/>`),
		g.Group(children),
	)
}