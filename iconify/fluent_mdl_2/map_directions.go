package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapDirections(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1024 0l1024 1024l-1024 1024L0 1024L1024 0zM768 1611l256 256l843-843l-843-843l-843 843l459 459V896h549L979 685l90-90l365 365l-365 365l-90-90l210-211H768v587z"/>`),
		g.Group(children),
	)
}