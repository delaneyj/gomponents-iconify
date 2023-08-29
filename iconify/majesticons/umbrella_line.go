package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4c-7.2 0-9 6-9 9h9m0-9c7.2 0 9 6 9 9h-9m0-9V3m0 10v5c0 1 .6 3 3 3s3-2 3-3"/>`),
		g.Group(children),
	)
}