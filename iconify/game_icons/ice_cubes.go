package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCubes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m90.01 41l46.09 452h239.8L422 41zM142 62.04l112.5 18.29l-3.9 24.17l104.2-21.92l7.7 36.42H394l-33.9 354H151.9L119 128.9l-1-9.9h14.7l7.8-48.07zm14.8 20.66l-12.5 77l77.1 12.5l12.5-76.99zm184 21.2L264.6 120l16 76.2l76.3-16zM219.5 242.6l-36.7 68.6l68.8 36.8l36.7-68.6z"/>`),
		g.Group(children),
	)
}