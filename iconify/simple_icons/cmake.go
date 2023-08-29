package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cmake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.769.066L.067 23.206l12.76-10.843zm11.438 23.868L7.471 17.587L0 23.934zm.793-.198L12.298.463l1.719 19.24zM12.893 12.959l-5.025 4.298l5.62 2.248z"/>`),
		g.Group(children),
	)
}