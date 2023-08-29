package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatbotSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3Z"/><path fill="currentColor" fill-rule="evenodd" d="M9 2H8V0H7v2H6a6 6 0 0 0 0 12h3c.13 0 .26-.004.389-.013l3.99.998a.5.5 0 0 0 .606-.606l-.577-2.309A6 6 0 0 0 9 2ZM5 6.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0ZM7.5 12a4.483 4.483 0 0 1-2.813-.987l.626-.78c.599.48 1.359.767 2.187.767c.828 0 1.588-.287 2.187-.767l.626.78A4.483 4.483 0 0 1 7.5 12Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}