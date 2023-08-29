package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 18.965a29.355 29.355 0 0 1 1.817-1.586C17.037 8.374 33.382 8.903 44 18.965"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 25.799c-7.732-7.732-20.268-7.732-28 0m22 6.515c-4.418-4.419-11.582-4.419-16 0"/><path fill="currentColor" fill-rule="evenodd" d="M24 40a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}