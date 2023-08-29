package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h16v3a7.998 7.998 0 0 1-4.124 7A7.998 7.998 0 0 1 20 19v3H4v-3a7.998 7.998 0 0 1 4.124-7A7.998 7.998 0 0 1 4 5V2Zm8 11a6 6 0 0 0-6 6v1h12v-1a6 6 0 0 0-6-6ZM6 4v1a6 6 0 1 0 12 0V4H6Z"/>`),
		g.Group(children),
	)
}