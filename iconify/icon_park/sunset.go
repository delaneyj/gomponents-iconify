package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sunset(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M19 8L33 34H5L19 8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M29 26L34 20L43 34H32"/><path stroke-linecap="round" stroke-linejoin="round" d="M10 41L38 41"/><circle cx="38" cy="10" r="3" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}