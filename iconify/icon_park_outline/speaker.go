package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M38 2H10a2 2 0 0 0-2 2v40a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Z"/><path d="M24 38a6 6 0 1 0 0-12a6 6 0 0 0 0 12Zm0-20a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g>`),
		g.Group(children),
	)
}