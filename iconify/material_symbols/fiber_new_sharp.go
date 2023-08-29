package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FiberNewSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 20V4h22v16H1Zm2.5-5h1.25v-3.5L7.3 15h1.2V9H7.25v3.5L4.75 9H3.5v6Zm6 0h4v-1.25H11v-1.1h2.5V11.4H11v-1.15h2.5V9h-4v6Zm5 0h6V9h-1.25v4.5h-1.1V10H16.9v3.5h-1.15V9H14.5v6Z"/>`),
		g.Group(children),
	)
}