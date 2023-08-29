package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stackoverflow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="m83.873 0l-9.166 6.81l34.08 45.801l9.164-6.81L83.873 0zM62.852 18.227l-7.309 8.771l43.885 36.531l7.308-8.775l-43.884-36.527zM46.04 42.033l-4.82 10.35l51.763 24.1l4.825-10.35l-51.768-24.1zm-9.863 25.924l-2.352 11.174l55.881 11.736l2.348-11.166l-55.877-11.744zm-26.13 14.555V128H112.45V82.512h-11.377l-.002 34.115H21.428V82.512h-11.38zM32.806 93.88v11.373h56.888V93.88H32.805z"/>`),
		g.Group(children),
	)
}