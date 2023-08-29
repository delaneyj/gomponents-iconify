package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" d="M38 20V8a2 2 0 0 0-2-2H12a2 2 0 0 0-2 2v12"/><rect width="36" height="22" x="6" y="20" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 34h15v8H20v-8Zm-8-8h3"/></g>`),
		g.Group(children),
	)
}