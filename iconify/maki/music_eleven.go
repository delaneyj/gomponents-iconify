package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9.5.5a.489.489 0 0 0-.22.05L3.5 2a.5.5 0 0 0-.5.5v4.59A1.5 1.5 0 1 0 4 8.5V5.38l5-1.25v1.46A1.5 1.5 0 1 0 10 7V1a.5.5 0 0 0-.5-.5zM4 4.38v-1.5l5-1.25v1.5L4 4.38z" fill="currentColor"/>`),
		g.Group(children),
	)
}