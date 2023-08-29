package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patreon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 .48v23.04h4.22V.48zm15.385 0c-4.764 0-8.641 3.88-8.641 8.65c0 4.755 3.877 8.623 8.641 8.623c4.75 0 8.615-3.868 8.615-8.623C24 4.36 20.136.48 15.385.48z"/>`),
		g.Group(children),
	)
}