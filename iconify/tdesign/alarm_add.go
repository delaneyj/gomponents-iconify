package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmAdd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.914 2.5L1.5 7.914L.086 6.5L5.5 1.086L6.914 2.5ZM18.5 1.086L23.914 6.5L22.5 7.914L17.086 2.5L18.5 1.086ZM12 5a8 8 0 1 0 0 16a8 8 0 0 0 0-16ZM2 13C2 7.477 6.477 3 12 3s10 4.477 10 10s-4.477 10-10 10S2 18.523 2 13Zm11-5v4h4v2h-4v4h-2v-4H7v-2h4V8h2Z"/>`),
		g.Group(children),
	)
}