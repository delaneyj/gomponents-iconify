package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m323.9 19.81l-55.2 55.15L285 91.24L272.2 104L256 87.73L19.81 323.9l45.57 45.6c28.5-14.6 56.22-11.7 72.52 4.6c16.3 16.3 19.2 44 4.6 72.5l45.6 45.6l236.1-236.1l-16.2-16.3l12.8-12.8l16.3 16.2l55.1-55.1l-45.6-45.6c-28.5 14.6-56.2 11.7-72.5-4.6c-16.3-16.3-19.2-44.02-4.6-72.52zm-16.2 93.99l33.9 34l-12.8 12.8l-33.9-34zM256 130.2L381.8 256L222.1 415.8L96.16 289.9L249.6 136.5zm0 25.4L121.6 289.9l100.5 100.5L356.4 256zm108.2 14.8l34 33.9l-12.8 12.8l-34-33.9z"/>`),
		g.Group(children),
	)
}