package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M24 19H40"/><path stroke-linecap="round" d="M24 10H40"/><path stroke-linecap="round" d="M8 38H40"/><path stroke-linecap="round" d="M8 28H40"/><path fill="#2F88FF" d="M8 10L16 15L8 20V10Z"/></g>`),
		g.Group(children),
	)
}