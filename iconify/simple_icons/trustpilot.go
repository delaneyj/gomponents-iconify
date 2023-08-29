package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trustpilot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.227 16.67l2.19 6.742l-7.413-5.388l5.223-1.354zM24 9.31h-9.165L12.005.589l-2.84 8.723L0 9.3l7.422 5.397l-2.84 8.714l7.422-5.388l4.583-3.326L24 9.311z"/>`),
		g.Group(children),
	)
}