package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alacritty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.065 0l-8.57 21.269H5.09L12 5.025l6.91 16.244h3.594L13.934 0zM12 9.935L9.702 15.5c1.475 4.54 1.475 4.54 2.298 8.5c.823-3.96.823-3.96 2.297-8.5L12 9.935z"/>`),
		g.Group(children),
	)
}