package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatPageBreak(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20H6v-2H4v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2h-2v2M14 2H6a2 2 0 0 0-2 2v8h2V4h8v4h4v4h2V8l-6-6m-3 14H8v-2h3v2m5 0h-3v-2h3v2M3 14h3v2H3v-2m18 2h-3v-2h3v2Z"/>`),
		g.Group(children),
	)
}