package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeIndesign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path d="M39 6H9a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 15v18m16-16v16"/><path stroke-linecap="round" stroke-linejoin="round" d="M28 33c-2.5 0-4-1.6-4-4s1.5-4 4-4h4v8h-4Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}