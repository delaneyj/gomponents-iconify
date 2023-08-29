package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path stroke="#fff" d="M24.5 14C21.4624 14 19 16.4624 19 19.5C19 22.5376 21.4624 25 24.5 25C27.5376 25 30 22.5376 30 19.5C30 16.4624 27.5376 14 24.5 14Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M19 29C19.8184 31.1932 21.5476 33 24.5 33C27.5376 33 30 30.3137 30 27V20"/></g>`),
		g.Group(children),
	)
}