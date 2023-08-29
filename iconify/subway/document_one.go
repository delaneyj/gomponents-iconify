package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M372.4 512L512 372.4H372.4V512zM0 0v512h349.1V349.1H512V0H0zm395.6 279.3H279.3v116.4h-46.5V279.3H116.4v-46.5h116.4V116.4h46.5v116.4h116.4v46.5z"/>`),
		g.Group(children),
	)
}