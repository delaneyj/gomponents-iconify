package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepBackwrad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.108 23.855L2.4 13.361v10.04a.599.599 0 0 1-.599.599H.597a.599.599 0 0 1-.599-.599V.599C-.002.268.266 0 .597 0h1.202c.331 0 .599.268.599.599v10.039L21.106.145a1.142 1.142 0 0 1 1.691 1.001v.024v-.001v21.68c0 .634-.511 1.149-1.143 1.155h-.001a1.108 1.108 0 0 1-.552-.152l.005.003z"/>`),
		g.Group(children),
	)
}