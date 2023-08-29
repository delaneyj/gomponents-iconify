package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 23h16V7H7v16ZM17 4V1h-3M1 14v3h3m-3-5V6v6ZM4 1H1v3m5-3h6h-6Z"/>`),
		g.Group(children),
	)
}