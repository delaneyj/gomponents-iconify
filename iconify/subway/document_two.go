package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M372.4 517.5L512 377.9H372.4v139.6zM0 5.5v512h349.1V354.6H512V5.5H0zm395.6 279.3H116.4v-46.5h279.3v46.5z"/>`),
		g.Group(children),
	)
}