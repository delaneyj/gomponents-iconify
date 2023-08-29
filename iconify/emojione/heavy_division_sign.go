package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyDivisionSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<g fill="#4d5357"><path d="M2 26h60v12H2z"/><circle cx="32" cy="9.5" r="7.5"/><circle cx="32" cy="54.5" r="7.5"/></g>`),
		g.Group(children),
	)
}