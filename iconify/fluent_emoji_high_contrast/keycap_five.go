package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M12.828 8.12a1.75 1.75 0 0 0-1.75 1.75v5.025c0 .967.784 1.75 1.75 1.75h2.719a2.11 2.11 0 1 1-1.582 3.505a1.75 1.75 0 0 0-2.624 2.316a5.61 5.61 0 1 0 4.206-9.321h-.969v-1.524h3.688a1.75 1.75 0 0 0 0-3.5h-5.438Z"/><path d="M6 1.12a5 5 0 0 0-5 5v20a5 5 0 0 0 5 5h20a5 5 0 0 0 5-5v-20a5 5 0 0 0-5-5H6Zm-3 5a3 3 0 0 1 3-3h20a3 3 0 0 1 3 3v20a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3v-20Z"/></g>`),
		g.Group(children),
	)
}