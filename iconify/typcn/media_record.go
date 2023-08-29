package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MediaRecord(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 12a5.985 5.985 0 0 0-1.757-4.243A5.985 5.985 0 0 0 12 6a5.985 5.985 0 0 0-4.242 1.757A5.982 5.982 0 0 0 6 12c0 1.656.672 3.156 1.758 4.242S10.344 18 12 18a5.982 5.982 0 0 0 4.243-1.758A5.985 5.985 0 0 0 18 12z"/>`),
		g.Group(children),
	)
}