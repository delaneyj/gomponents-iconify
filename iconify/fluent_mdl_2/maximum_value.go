package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MaximumValue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 512h128v896H0V512zm2048 0v896h-128V512h128zm-531 131l318 317l-318 317l-90-90l162-163H475l162 163l-90 90l-317-317l317-317l90 90l-162 163h1114l-162-163l90-90z"/>`),
		g.Group(children),
	)
}