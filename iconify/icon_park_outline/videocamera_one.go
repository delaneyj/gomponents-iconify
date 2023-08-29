package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideocameraOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="31" height="20" x="4" y="21" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M20 27h9v8h-9z"/><circle cx="27" cy="13" r="5" stroke-linecap="round" stroke-linejoin="round"/><circle cx="13" cy="13" r="5" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="m35 35l9 4V23l-9 4"/></g>`),
		g.Group(children),
	)
}