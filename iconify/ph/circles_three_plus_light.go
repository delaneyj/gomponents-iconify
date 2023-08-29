package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CirclesThreePlusLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M80 42a38 38 0 1 0 38 38a38 38 0 0 0-38-38Zm0 64a26 26 0 1 1 26-26a26 26 0 0 1-26 26Zm96 12a38 38 0 1 0-38-38a38 38 0 0 0 38 38Zm0-64a26 26 0 1 1-26 26a26 26 0 0 1 26-26Zm-96 84a38 38 0 1 0 38 38a38 38 0 0 0-38-38Zm0 64a26 26 0 1 1 26-26a26 26 0 0 1-26 26Zm134-26a6 6 0 0 1-6 6h-26v26a6 6 0 0 1-12 0v-26h-26a6 6 0 0 1 0-12h26v-26a6 6 0 0 1 12 0v26h26a6 6 0 0 1 6 6Z"/>`),
		g.Group(children),
	)
}