package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoveAndHelp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M15 7C8.92487 7 4 11.9249 4 18C4 29 17 39 24 41.3262C31 39 44 29 44 18C44 11.9249 39.0751 7 33 7C29.2797 7 25.9907 8.8469 24 11.6738C22.0093 8.8469 18.7203 7 15 7Z"/><path stroke="#fff" d="M18 24H30"/><path stroke="#fff" d="M24 18V30"/></g>`),
		g.Group(children),
	)
}