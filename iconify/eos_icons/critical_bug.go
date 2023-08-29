package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CriticalBug(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 8h-2.8a6 6 0 0 0-1.82-2L17 4.41L15.6 3l-2.17 2.17A6.07 6.07 0 0 0 12 5a5.92 5.92 0 0 0-1.41.17L8.42 3L7 4.41L8.63 6a6.06 6.06 0 0 0-1.81 2H4v2h2.1a6.64 6.64 0 0 0-.1 1v1H4v2h2v1a6.64 6.64 0 0 0 .09 1H4v2h2.82a6 6 0 0 0 10.38 0H20v-2h-2.08a6.64 6.64 0 0 0 .08-1v-1h2v-2h-2v-1a6.64 6.64 0 0 0-.09-1H20Zm-7 10h-2v-2h2Zm0-4h-2V8h2Z"/>`),
		g.Group(children),
	)
}