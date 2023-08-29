package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridKanbanSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V4Zm4 0H3v1h3V4Zm3 0a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1V4Zm4 0h-3v1h3V4ZM3 7a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H3Zm0 1h3v1H3V8Zm6 0a1 1 0 0 1 1-1h3a1 1 0 0 1 1 1v1a1 1 0 0 1-1 1h-3a1 1 0 0 1-1-1V8Zm4 0h-3v1h3V8ZM3 11a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1H3Zm0 1h3v1H3v-1Z"/>`),
		g.Group(children),
	)
}