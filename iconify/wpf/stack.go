package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M5 0C3.344 0 2.625 1.422 2 3L.062 8.25C.346 8.094.654 8 1 8h24c.346 0 .654.094.938.25L24 3c-.703-1.531-1.344-3-3-3H5zM1 9c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h24c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1H1zm21.5 1a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 22.5 10zM1 15c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h24c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1H1zm21.5 1a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 22.5 16zM1 21c-.551 0-1 .449-1 1v3c0 .551.449 1 1 1h24c.551 0 1-.449 1-1v-3c0-.551-.449-1-1-1H1zm21.5 1a1.5 1.5 0 1 1-.001 3.001A1.5 1.5 0 0 1 22.5 22z"/>`),
		g.Group(children),
	)
}