package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M9 10.5a.5.5 0 0 1-.5.5a.49.49 0 0 1-.43-.27L7.69 10H3.31l-.37.74A.5.5 0 0 1 2 10.5a.49.49 0 0 1 .07-.24l1-2A.49.49 0 0 1 3.5 8a.5.5 0 0 1 .5.7l-.19.3h3.38L7 8.7a.5.5 0 0 1 .9-.43l1 2a.49.49 0 0 1 .1.23zM8 0H3a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1zM3.5 6a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zm0-2a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 1 .5-.5h1.79v3H3.5zm4 2a.5.5 0 1 1 0-1a.5.5 0 0 1 0 1zM8 3.5a.5.5 0 0 1-.5.5H5.69V1H7.5a.5.5 0 0 1 .5.5v2z" fill="currentColor"/>`),
		g.Group(children),
	)
}