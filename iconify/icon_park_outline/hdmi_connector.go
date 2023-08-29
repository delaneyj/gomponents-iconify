package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HdmiConnector(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><circle cx="24" cy="24" r="20" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M29 43a5 5 0 0 0-10 0"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M15.5 42.11A19.923 19.923 0 0 0 24 44c3.04 0 5.92-.678 8.5-1.89"/><circle cx="15" cy="15" r="3" fill="currentColor"/><circle cx="11" cy="23" r="3" fill="currentColor"/><circle cx="24" cy="11" r="3" fill="currentColor"/><circle cx="33" cy="15" r="3" fill="currentColor"/><circle cx="37" cy="23" r="3" fill="currentColor"/></g>`),
		g.Group(children),
	)
}