package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.17 10.17a1 1 0 0 0 .71-.29a1 1 0 0 0 0-1.42L5.41 4H7a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v4a1 1 0 0 0 2 0V5.41l4.46 4.47a1 1 0 0 0 .71.29Zm6.37-1.71a1 1 0 0 0-1.42 0l-5.66 5.66a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29l5.66-5.66a1 1 0 0 0 0-1.42ZM21 16a1 1 0 0 0-1 1v1.59l-4.46-4.47a1 1 0 1 0-1.42 1.42L18.59 20H17a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-4a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}