package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12.5 10H11V6h1.5A2.5 2.5 0 1 0 10 3.5V5H6V3.5A2.5 2.5 0 1 0 3.5 6H5v4H3.5A2.5 2.5 0 1 0 6 12.5V11h4v1.5a2.5 2.5 0 1 0 2.5-2.5zM11 3.5A1.5 1.5 0 1 1 12.5 5H11V3.5zm-6 9A1.5 1.5 0 1 1 3.5 11H5v1.5zM5 5H3.5A1.5 1.5 0 1 1 5 3.5V5zm5 5H6V6h4v4zm2.5 4a1.5 1.5 0 0 1-1.5-1.5V11h1.5a1.5 1.5 0 0 1 0 3z"/>`),
		g.Group(children),
	)
}