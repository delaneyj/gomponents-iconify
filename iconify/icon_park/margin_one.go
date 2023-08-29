package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarginOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path stroke="#fff" stroke-linecap="round" d="M6 34H42"/><path stroke="#fff" stroke-linecap="round" d="M6 14H42"/><path stroke="#000" stroke-linecap="round" d="M6 11V17"/><path stroke="#000" stroke-linecap="round" d="M6 31V37"/><path stroke="#000" stroke-linecap="round" d="M42 31V37"/><path stroke="#000" stroke-linecap="round" d="M42 11V17"/></g>`),
		g.Group(children),
	)
}