package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileAudioFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2H5Zm9 2.414L17.586 8H14V4.414ZM11 15v-3a.997.997 0 0 1 1-1a.997.997 0 0 1 .707.293l2 2a1 1 0 0 1-1.414 1.414L13 14.414V17a2 2 0 1 1-2-2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}