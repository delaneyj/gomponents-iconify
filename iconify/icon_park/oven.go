package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="40" height="28" x="4" y="8" fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="2"/><rect width="16" height="12" x="12" y="16" fill="#43CCF8" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" rx="1"/><circle cx="37" cy="15" r="2" fill="#fff"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 22H38"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 29H38"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M12 36L12 42"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M36 36L36 42"/></g>`),
		g.Group(children),
	)
}