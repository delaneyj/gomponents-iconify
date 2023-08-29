package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h12v4H3V4Zm18 4h-4V4h4v4ZM3 10h18v4H3v-4Zm8 6H3v4h8v-4Zm2 0v4h8v-4h-8Z"/>`),
		g.Group(children),
	)
}