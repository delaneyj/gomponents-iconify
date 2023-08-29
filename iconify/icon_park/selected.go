package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Selected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M12 4H4V12H12V4Z"/><path fill="#2F88FF" d="M44 36H36V44H44V36Z"/><path fill="#2F88FF" d="M12 36H4V44H12V36Z"/><path fill="#2F88FF" d="M44 4H36V12H44V4Z"/><path stroke-linecap="round" d="M8 36V12"/><path stroke-linecap="round" d="M40 36V12"/><path stroke-linecap="round" d="M12 8H36"/><path stroke-linecap="round" d="M12 40H36"/></g>`),
		g.Group(children),
	)
}