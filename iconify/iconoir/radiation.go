package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 20.662a9.955 9.955 0 0 1-5 1.337a9.954 9.954 0 0 1-5-1.337L10 16s1 .4 2 .4s2-.4 2-.4l3 4.662Zm-.002-17.323A9.954 9.954 0 0 1 20.656 7a9.954 9.954 0 0 1 1.342 5l-5.537-.268s-.154-1.066-.654-1.932c-.5-.866-1.346-1.532-1.346-1.532l2.537-4.93ZM1.998 12A9.954 9.954 0 0 1 3.34 7a9.954 9.954 0 0 1 3.658-3.66l2.537 4.928S8.69 8.934 8.19 9.8s-.654 1.932-.654 1.932L1.998 12ZM12 14a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/>`),
		g.Group(children),
	)
}