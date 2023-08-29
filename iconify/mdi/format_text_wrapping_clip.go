package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextWrappingClip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21H5V3h2v18M17 3v8H9v2h8v8h2V3h-2Z"/>`),
		g.Group(children),
	)
}