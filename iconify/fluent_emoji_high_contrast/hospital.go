package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M12.5 3.5a1 1 0 0 0-1 1v1h-1a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0v-1h1a1 1 0 1 0 0-2h-1v-1a1 1 0 0 0-1-1Z"/><path d="M16 1a2 2 0 0 1 2 2h4a2 2 0 0 1 2 2v14h2a2 2 0 0 1 2 2v9h-2v-1.5a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.5.5V30h-2.5v-1.5a.5.5 0 0 0-.5-.5h-2v2h-1.5v-2h-2a.5.5 0 0 0-.5.5V30H5V5a2 2 0 0 1 2-2a2 2 0 0 1 2-2h7ZM8 3v7a1 1 0 0 0 1 1h7a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1Zm3.75 18a.75.75 0 0 0 0 1.5h13.5a.75.75 0 0 0 0-1.5h-13.5Zm0 2.5a.75.75 0 0 0 0 1.5h13.5a.75.75 0 0 0 0-1.5h-13.5Z"/></g>`),
		g.Group(children),
	)
}