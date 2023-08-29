package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderBigender(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11a4 4 0 1 0 8 0a4 4 0 1 0-8 0m12-8l-5 5m1-5h4v4m-8 9v6m-3-3h6"/>`),
		g.Group(children),
	)
}