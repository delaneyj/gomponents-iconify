package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DuplicateAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M14.5 16.5v-8a2 2 0 0 0-2-2h-8a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2z"/><path d="M14.5 14.5h2a2 2 0 0 0 2-2V4.503a2 2 0 0 0-2-2h-.003l-8 .014a2 2 0 0 0-1.997 2V6.5m2 3v6m3-3h-6"/></g>`),
		g.Group(children),
	)
}