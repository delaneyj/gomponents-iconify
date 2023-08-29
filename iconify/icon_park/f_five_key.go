package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FFiveKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" d="M34 16H26V22.9185C26 22.9185 27.8333 21.5 30.5 21.5C33.1667 21.5 34 23.6585 34 27C34 30.3415 32.5 32 29.5556 32C26.8889 32 26 30.315 26 28.6365"/><path stroke="#fff" d="M21 16H14V32"/><path stroke="#fff" d="M14 24H21"/></g>`),
		g.Group(children),
	)
}