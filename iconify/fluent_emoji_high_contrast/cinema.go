package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cinema(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M15.5 13a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7ZM12 16.5a2.5 2.5 0 0 1 2.5-2.5h8a2.5 2.5 0 0 1 2.5 2.5v5a2.5 2.5 0 0 1-2.5 2.5h-8a2.5 2.5 0 0 1-2.5-2.5v-5Zm-4.953-1.637v8.326a.75.75 0 0 0 1.28.53l2.469-2.468a.75.75 0 0 0 .22-.53v-3.409a.75.75 0 0 0-.222-.532l-2.469-2.45a.75.75 0 0 0-1.278.533ZM25 10.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/><path d="M6 1a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5V6a5 5 0 0 0-5-5H6ZM3 6a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V6Z"/></g>`),
		g.Group(children),
	)
}