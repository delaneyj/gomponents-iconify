package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Document(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M372.4 512L512 372.4H372.4V512zM0 0v512h349.1V349.1H512V0H0z"/>`),
		g.Group(children),
	)
}