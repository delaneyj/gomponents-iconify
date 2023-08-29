package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CutleryF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.955 10.5l-8.298 8.3a2 2 0 1 1-2.829-2.828l8.299-8.299c-.173-1.518.515-3.343 1.954-4.783C13.23.741 16.237.266 17.8 1.83c1.562 1.562 1.087 4.57-1.06 6.717c-1.44 1.44-3.266 2.128-4.784 1.955z"/>`),
		g.Group(children),
	)
}