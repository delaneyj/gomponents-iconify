package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Door(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M22.271 14.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0ZM10 5a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1V5Zm1 .5v6a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-9a.5.5 0 0 0-.5.5ZM11 16a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1H11Zm.5 1h9a.5.5 0 0 1 .5.5v2a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5ZM10 24a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v3a1 1 0 0 1-1 1H11a1 1 0 0 1-1-1v-3Zm1 .5v2a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5h-9a.5.5 0 0 0-.5.5Z"/><path d="M7 4a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v24a3 3 0 0 1-3 3H10a3 3 0 0 1-3-3V4Zm3-1a1 1 0 0 0-1 1v24a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H10Z"/></g>`),
		g.Group(children),
	)
}