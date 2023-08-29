package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 44H44"/><path fill="#2F88FF" stroke-linejoin="round" d="M8 8.36364L24 4L40 8.36364V14H8V8.36364Z"/><path stroke-linecap="round" d="M10 14V38"/><path stroke-linecap="round" d="M17 14V38"/><path stroke-linecap="round" d="M24 14V38"/><path stroke-linecap="round" d="M31 14V38"/><path stroke-linecap="round" d="M38 14V38"/><rect width="34" height="6" x="7" y="38" stroke-linejoin="round"/></g>`),
		g.Group(children),
	)
}