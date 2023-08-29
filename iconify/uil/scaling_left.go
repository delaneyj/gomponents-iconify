package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScalingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16a1 1 0 0 0-1 1v1.59L13.41 12l2.13-2.12a1 1 0 0 0-1.42-1.42L12 10.59L5.41 4H7a1 1 0 0 0 0-2H3a1 1 0 0 0-.38.08a1 1 0 0 0-.54.54A1 1 0 0 0 2 3v4a1 1 0 0 0 2 0V5.41L10.59 12l-2.13 2.12a1 1 0 0 0 0 1.42a1 1 0 0 0 .71.29a1 1 0 0 0 .71-.29L12 13.41L18.59 20H17a1 1 0 0 0 0 2h4a1 1 0 0 0 .38-.08a1 1 0 0 0 .54-.54A1 1 0 0 0 22 21v-4a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}