package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4V6Z"/><path fill="#fff" d="M8 9.5A1.5 1.5 0 0 1 9.5 8h13A1.5 1.5 0 0 1 24 9.5v13a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 8 22.5v-13Z"/></g>`),
		g.Group(children),
	)
}