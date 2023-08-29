package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 0v16h16V0H0zm5 15H1v-4h4v4zm0-5H1V6h4v4zm0-5H1V1h4v4zm5 10H6v-4h4v4zm0-5H6V6h4v4zm0-5H6V1h4v4zm5 10h-4v-4h4v4zm0-5h-4V6h4v4zm0-5h-4V1h4v4z"/>`),
		g.Group(children),
	)
}