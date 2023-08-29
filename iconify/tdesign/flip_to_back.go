package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipToBack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 3h3.5v2H9v1.5H7V3Zm5.5 0h3v2h-3V3Zm5 0H21v3.5h-2V5h-1.5V3ZM5 7.5V19h11.5v2H3V7.5h2Zm4 1v3H7v-3h2Zm12 0v3h-2v-3h2Zm-12 5V15h1.5v2H7v-3.5h2Zm12 0V17h-3.5v-2H19v-1.5h2ZM12.5 15h3v2h-3v-2Z"/>`),
		g.Group(children),
	)
}