package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pacman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.636 5.636a9 9 0 0 1 13.397.747L13.414 12l5.619 5.617A9 9 0 1 1 5.636 5.636z"/><circle cx="11.5" cy="7.5" r="1" fill="currentColor"/></g>`),
		g.Group(children),
	)
}