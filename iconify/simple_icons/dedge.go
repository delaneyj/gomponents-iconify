package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dedge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.986 0v8.338C16.09 2.93 7.61 2.8 3.74 8.733C-.523 15.27 4.191 23.99 11.996 24h.001c5.447-.003 9.872-4.43 9.87-9.877V0Zm-7.99 6.14a8.004 8.004 0 0 1 7.99 7.988a7.986 7.986 0 0 1-4.93 7.381a7.986 7.986 0 0 1-8.707-1.73a7.985 7.985 0 0 1-1.733-8.707a7.986 7.986 0 0 1 7.38-4.932Z"/>`),
		g.Group(children),
	)
}