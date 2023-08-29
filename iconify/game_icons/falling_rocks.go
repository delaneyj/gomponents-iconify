package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FallingRocks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m156.115 34.48l-36.94.586l10.02 28.995l27.258-2.324zM18 37.727V494h297.682L304 448l-64-64l-48-160l-88.055-80L96 64zM219.076 98.58L160 112l16.38 30.342l17.65 5.394l-1.223 19.672L244.664 144zm34.656 60.99l-19.56 21.05l27.508 12.61l12.855-20.804zm50.907 90.002l-34.38 25.012l-11.766 58.662l100.53 22.97l-4.163-98.218zm107.475 62.373l-20.886 3.465l-15.988 20.033l4.873 21.23l40.848-15.55zm-43.62 95.996l-44.32 5.87l6.858 29.422l36.875-1.396zm37.573 59.982l-32.103 5.375L340.113 494h79.508z"/>`),
		g.Group(children),
	)
}