package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CitySeven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 2H2v20h20V10h-8V2Zm-2 8H8v10H4V4h8v6Zm-2 10v-8h10v8h-4v-4h-2v4h-4Z"/>`),
		g.Group(children),
	)
}