package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Electrical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m22 3.594l-4 3.969l-2.28-2.282l-1.44 1.44l.75.75l-5.124 5.124a3.124 3.124 0 0 0 0 4.406l1.844 1.844l-7.47 7.437l1.44 1.44l7.436-7.47L15 22.094a3.124 3.124 0 0 0 4.406 0l5.125-5.125l.75.75l1.44-1.44L24.437 14l3.968-4L27 8.594l-4 3.97L19.437 9l3.97-4L22 3.594zm-5.563 5.28l6.688 6.688L18 20.689c-.388.387-1.206.387-1.594 0l-5.093-5.094c-.388-.388-.388-1.206 0-1.594l5.124-5.125z"/>`),
		g.Group(children),
	)
}