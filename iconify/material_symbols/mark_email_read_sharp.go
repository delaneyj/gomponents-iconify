package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkEmailReadSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.95 22l-4.25-4.25l1.4-1.4l2.85 2.85l5.65-5.65l1.4 1.4L15.95 22ZM2 20V4h20v6.35l-6.025 6.025L13.1 13.5l-4.225 4.225L11.15 20H2Zm10-7l8-5V6l-8 5l-8-5v2l8 5Z"/>`),
		g.Group(children),
	)
}