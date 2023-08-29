package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DownloadDoneOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20v-2h14v2H5Zm4.55-4l-5.675-5.675L5.3 8.9l4.25 4.25L18.7 4l1.425 1.425L9.55 16Z"/>`),
		g.Group(children),
	)
}