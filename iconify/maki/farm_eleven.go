package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FarmEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6 5L4 4L2 5L1 7v3h1.5V8h3v2H7V7zM5 7H3V5.5h2V7z" fill="currentColor"/><path d="M10 2a1 1 0 0 0-2 0v8h2V2z" fill="currentColor"/>`),
		g.Group(children),
	)
}