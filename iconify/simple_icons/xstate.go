package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Xstate(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.891 0h6.023l-6.085 10.563A10.653 10.653 0 0 1 15.891 0zm6.055 23.999L8.078.001H2.055l6.919 12.015L2.055 24h6.023L12 17.236L15.892 24z"/>`),
		g.Group(children),
	)
}