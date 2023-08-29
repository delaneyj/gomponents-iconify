package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17.772 9.5a1.75 1.75 0 1 0-3.5 0v3.482l-3.026-1.747a1.75 1.75 0 1 0-1.75 3.03L12.5 16l-3.004 1.735a1.75 1.75 0 1 0 1.75 3.03l3.026-1.747V22.5a1.75 1.75 0 1 0 3.5 0v-3.456l2.982 1.722a1.75 1.75 0 0 0 1.75-3.031L19.5 16l3.004-1.734a1.75 1.75 0 0 0-1.75-3.031l-2.982 1.721V9.5Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}