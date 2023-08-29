package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourBox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 512h232.7V279.3H0V512zm0-279.3h232.7V0H0v232.7zM279.3 512H512V279.3H279.3V512zm0-512v232.7H512V0H279.3z"/>`),
		g.Group(children),
	)
}