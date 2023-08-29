package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><path fill="#00A6ED" d="M2 6.12a4 4 0 0 1 4-4h20a4 4 0 0 1 4 4v20a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4v-20Z"/><path fill="#fff" d="M11.078 9.87c0-.966.784-1.75 1.75-1.75h5.438a1.75 1.75 0 0 1 0 3.5h-3.688v1.525h.969a5.61 5.61 0 1 1-4.206 9.321a1.75 1.75 0 0 1 2.624-2.316a2.11 2.11 0 1 0 1.582-3.505h-2.719a1.75 1.75 0 0 1-1.75-1.75V9.87Z"/></g>`),
		g.Group(children),
	)
}