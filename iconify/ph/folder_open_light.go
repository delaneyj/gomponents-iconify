package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpenLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M243.36 111.81A14 14 0 0 0 232 106h-18V88a14 14 0 0 0-14-14h-69.33a2 2 0 0 1-1.2-.4l-27.73-20.8a14.06 14.06 0 0 0-8.4-2.8H40a14 14 0 0 0-14 14v144a6 6 0 0 0 6 6h179.1a6 6 0 0 0 5.69-4.1l28.49-85.47a14 14 0 0 0-1.92-12.62ZM40 62h53.34a2 2 0 0 1 1.2.4l27.73 20.8a14.06 14.06 0 0 0 8.4 2.8H200a2 2 0 0 1 2 2v18H69.77a14 14 0 0 0-13.28 9.57L38 171V64a2 2 0 0 1 2-2Zm193.9 58.63L206.78 202H40.33l27.54-82.63a2 2 0 0 1 1.9-1.37H232a2 2 0 0 1 1.9 2.63Z"/>`),
		g.Group(children),
	)
}