package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 6v29"/><path d="M10 36.04A5.04 5.04 0 0 1 15.04 31H24v5.96A5.04 5.04 0 0 1 18.96 42h-3.92A5.04 5.04 0 0 1 10 36.96v-.92Z"/><path stroke-linecap="round" d="m24 14.066l12.883 3.056V9.013L24 6v8.066Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}