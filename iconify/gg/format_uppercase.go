package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatUppercase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 9h-3v8H8V9H5V7h8v2Zm5 4h-2v4h-2v-4h-2v-2h6v2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}