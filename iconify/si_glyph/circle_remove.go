package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleRemove(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.021 1.097C3.625 1.097.063 4.655.063 9.04c0 4.388 3.562 7.945 7.958 7.945c4.395 0 7.958-3.558 7.958-7.945c0-4.386-3.564-7.943-7.958-7.943zM10.271 10H5.729C4.772 10 4 10.05 4 9c0-1.053.772-1 1.728-1h4.544C11.228 8 12 7.946 12 9c0 1.051-.772 1-1.728 1z"/>`),
		g.Group(children),
	)
}