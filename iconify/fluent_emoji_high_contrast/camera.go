package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18 21a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/><path d="M16 27a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-2a4 4 0 1 1 0-8a4 4 0 0 1 0 8Z"/><path d="M5.5 10a.5.5 0 0 0-.5.5v.5H4a3 3 0 0 0-3 3v14a3 3 0 0 0 3 3h24a3 3 0 0 0 3-3V14a3 3 0 0 0-3-3h-1v-.5a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.5.5v.5H9v-.5a.5.5 0 0 0-.5-.5h-3ZM3 14a1 1 0 0 1 1-1h24a1 1 0 0 1 1 1v1h-9.392A6.996 6.996 0 0 1 23 21a6.996 6.996 0 0 1-3.392 6H29v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1v-1h9.392A6.996 6.996 0 0 1 9 21a6.996 6.996 0 0 1 3.392-6H3v-1Z"/></g>`),
		g.Group(children),
	)
}