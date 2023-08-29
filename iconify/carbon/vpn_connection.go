package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VpnConnection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 7a1.992 1.992 0 0 0-1.723 1H22a6 6 0 0 0-12 0H5.723a2 2 0 1 0 0 2H10v4H8a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V16a2 2 0 0 0-2-2h-2v-4h4.277A1.997 1.997 0 1 0 28 7ZM12 8a4 4 0 0 1 8 0v6h-8v-4h5V8Zm12 8v12H8V16Z"/>`),
		g.Group(children),
	)
}