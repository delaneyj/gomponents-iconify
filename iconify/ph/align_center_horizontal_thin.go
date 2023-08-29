package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 140h-76v-24h52a12 12 0 0 0 12-12V64a12 12 0 0 0-12-12h-52V32a4 4 0 0 0-8 0v20H72a12 12 0 0 0-12 12v40a12 12 0 0 0 12 12h52v24H48a12 12 0 0 0-12 12v40a12 12 0 0 0 12 12h76v20a4 4 0 0 0 8 0v-20h76a12 12 0 0 0 12-12v-40a12 12 0 0 0-12-12ZM68 104V64a4 4 0 0 1 4-4h112a4 4 0 0 1 4 4v40a4 4 0 0 1-4 4H72a4 4 0 0 1-4-4Zm144 88a4 4 0 0 1-4 4H48a4 4 0 0 1-4-4v-40a4 4 0 0 1 4-4h160a4 4 0 0 1 4 4Z"/>`),
		g.Group(children),
	)
}