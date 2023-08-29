package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExitOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 20v-6h6v2h-4v4h-2Zm-6 0v-4H4v-2h6v6H8Zm12-10h-6V4h2v4h4v2Zm-10 0H4V8h4V4h2v6Z"/>`),
		g.Group(children),
	)
}