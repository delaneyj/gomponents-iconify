package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InkEraserOffOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.775 22.625L15.05 17.9L13 20H4.75L1.2 16.45l6.1-6.3l-5.925-5.925L2.8 2.8l18.4 18.4l-1.425 1.425ZM5.6 18h6.55l1.475-1.525L8.7 11.55L4 16.4L5.6 18Zm12.275-2.975L16.45 13.6L20 9.95L15.05 5L11.5 8.65l-1.4-1.4l4.9-5.1l7.8 7.8l-4.925 5.075Zm-3.9-3.9ZM11.175 14Z"/>`),
		g.Group(children),
	)
}