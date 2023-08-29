package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><rect width="26" height="40" x="11" y="4" rx="2"/><circle cx="24" cy="34" r="4"/><path stroke-linecap="round" stroke-linejoin="round" d="M18 10h12v8H18zm5 14h2"/></g>`),
		g.Group(children),
	)
}