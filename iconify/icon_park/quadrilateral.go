package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Quadrilateral(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#2F88FF" stroke="#000" stroke-width="4" d="M28.038 8H7C5.34315 8 4 9.34315 4 11V37C4 38.6569 5.34315 40 7 40H39.413C41.5764 40 43.0286 37.7796 42.1614 35.7975L30.7864 9.79754C30.3087 8.70558 29.2298 8 28.038 8Z"/>`),
		g.Group(children),
	)
}