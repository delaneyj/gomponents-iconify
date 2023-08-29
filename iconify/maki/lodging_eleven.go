package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LodgingEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M1.5 2a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 1 0V8h7v.5a.5.5 0 0 0 1 0v-1a.5.5 0 0 0-.5-.5H2V2.5a.5.5 0 0 0-.5-.5zm2 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2zM6 3a1 1 0 0 0-1 1v1H3a.5.5 0 0 0 0 1h7V5a2 2 0 0 0-2-2H6z" fill="currentColor"/>`),
		g.Group(children),
	)
}