package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerprintOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10v9h-2v-9a8 8 0 1 0-16 0v9H2v-9Zm4.5 0a5.5 5.5 0 1 1 11 0v9h-2v-9a3.5 3.5 0 1 0-7 0v9h-2v-9Zm6.5-1v10h-2V11h2Z"/>`),
		g.Group(children),
	)
}