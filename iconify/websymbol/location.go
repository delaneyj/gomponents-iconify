package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Location(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M586 294q0 7-1.5 21T567 386.5t-43.5 130t-88 202T293 1001q-82-153-142.5-282t-88-204t-43-128T2 314l-2-20Q0 173 86 87T293 1t207 86t86 207zm-207 0q0-36-25-61t-61-25t-61 25t-25 61t25 61t61 25t61-25t25-61z"/>`),
		g.Group(children),
	)
}