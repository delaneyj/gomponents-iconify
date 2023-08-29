package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneSignalWifiConnectedNoInternetFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 8.98A16.88 16.88 0 0 0 12 4C7.31 4 3.07 5.9 0 8.98L12 21v-9h8.99L24 8.98zM19.59 14l-2.09 2.09l-.3-.3L15.41 14L14 15.41l1.79 1.79l.3.3L14 19.59L15.41 21l2.09-2.08L19.59 21L21 19.59l-2.08-2.09L21 15.41L19.59 14z"/>`),
		g.Group(children),
	)
}