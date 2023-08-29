package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVerticalVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 4l-4 4h3v8a2 2 0 0 1-2 2a2 2 0 0 1-2-2V8a4 4 0 0 0-4-4a4 4 0 0 0-4 4v8H2l4 4l4-4H7V8a2 2 0 0 1 2-2a2 2 0 0 1 2 2v8a4 4 0 0 0 4 4a4 4 0 0 0 4-4V8h3l-4-4Z"/>`),
		g.Group(children),
	)
}