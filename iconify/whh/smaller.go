package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smaller(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 855v129q0 27-18.5 36t-45.5-6L64 684q-27-15-45.5-44T0 584V439q0-26 18.5-55.5T64 340L960 10q27-15 45.5-6.5T1024 39v129q0 27-18.5 56.5T960 268L299 512l661 244q27 14 45.5 43t18.5 56z"/>`),
		g.Group(children),
	)
}