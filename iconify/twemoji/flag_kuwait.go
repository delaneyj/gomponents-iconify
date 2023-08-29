package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagKuwait(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#007A3D" d="M32 5H4a4 4 0 0 0-4 4v4.5h36V9a4 4 0 0 0-4-4z"/><path fill="#CE1126" d="M0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-4.5H0V27z"/><path fill="#EEE" d="M0 13.5h36v9H0z"/><path fill="#141414" d="M1.205 6.138A3.993 3.993 0 0 0 0 9v18c0 1.122.462 2.135 1.205 2.862L9 22.5v-9L1.205 6.138z"/>`),
		g.Group(children),
	)
}