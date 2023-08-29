package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><rect width="24" height="36" x="12" y="6" fill="#2F88FF" rx="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12H6"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 20H6"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 28H6"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 36H6"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 12H36"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 20H36"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 28H36"/><path stroke-linecap="round" stroke-linejoin="round" d="M42 36H36"/></g>`),
		g.Group(children),
	)
}