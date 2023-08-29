package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Monster(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0v24h5.42V12.39l6.58 5.8l6.58-5.8V24H24V0L12 11.23L0 0Z"/>`),
		g.Group(children),
	)
}