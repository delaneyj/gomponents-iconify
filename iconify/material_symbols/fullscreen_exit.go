package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullscreenExit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 19v-3H5v-2h5v5H8Zm6 0v-5h5v2h-3v3h-2Zm-9-9V8h3V5h2v5H5Zm9 0V5h2v3h3v2h-5Z"/>`),
		g.Group(children),
	)
}