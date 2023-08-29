package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M4 44H44"/><path fill="#2F88FF" d="M8 7.27273C8 7.27273 15 4 24 4C33 4 40 7.27273 40 7.27273V13H8V7.27273Z"/><path stroke-linecap="round" d="M10 13V38"/><path stroke-linecap="round" d="M17 13V38"/><path stroke-linecap="round" d="M24 13V38"/><path stroke-linecap="round" d="M31 13V38"/><path stroke-linecap="round" d="M38 13V38"/><rect width="34" height="6" x="7" y="38"/></g>`),
		g.Group(children),
	)
}