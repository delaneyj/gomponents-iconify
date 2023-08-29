package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dialpad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 3h3v3h-3V3Zm0 5h3v3h-3V8Zm0 5v3h3v-3h-3Zm5-10h3v3h-3V3Zm0 5v3h3V8h-3Zm0 5h3v3h-3v-3Zm0 5v3h3v-3h-3Zm5-15h3v3h-3V3Zm0 5v3h3V8h-3Zm0 5h3v3h-3v-3Z"/>`),
		g.Group(children),
	)
}