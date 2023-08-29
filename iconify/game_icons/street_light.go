package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StreetLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m262.5 33l-10 30h87l-10-30h-67zM160 39c-26 0-45.2 9.12-56.9 23.24C91.32 76.35 87 94.5 87 112v263h18V112c0-14.5 3.7-28.35 11.9-38.24C125.2 63.88 138 57 160 57h75.5l6-18H160zm79.3 42l-43.7 42.5l41.9-19.1l-34.5 86.9l62.6-58.3l-1.1 91.7l31.8-101.2l70.5 117.6l-31.3-130.9l61.5 36.1L349.5 81H239.3zM72.55 393l-5.08 100h57.03l-5.1-100H72.55z"/>`),
		g.Group(children),
	)
}