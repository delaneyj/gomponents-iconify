package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicMenu(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M29 6v29"/><path d="M15 36.04A5.04 5.04 0 0 1 20.04 31H29v5.96A5.04 5.04 0 0 1 23.96 42h-3.92A5.04 5.04 0 0 1 15 36.96v-.92Z"/><path stroke-linecap="round" d="m29 14.066l12.883 3.056V9.013L29 6v8.066Z" clip-rule="evenodd"/><path stroke-linecap="round" d="M6 8h14M6 16h14M6 24h10"/></g>`),
		g.Group(children),
	)
}