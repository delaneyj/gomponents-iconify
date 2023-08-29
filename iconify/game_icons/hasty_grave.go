package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HastyGrave(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m416.62 213.555l-282.14-28.67l6.78-66.7l282.14 28.67zM79.38 492.185h353.24s-72.39-55.05-177-55.05c-52.38 0-112.77 13.76-176.24 55.05zm205.27-69.68l20.16-204.23l-66.66-6.77l-20.92 212a313.48 313.48 0 0 1 38.35-2.36c9.99-.01 19.69.48 29.07 1.36zm39.1-396.11l-66.68-6.58l-9.2 93.17l66.66 6.77z"/>`),
		g.Group(children),
	)
}