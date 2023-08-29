package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JewelryStoreEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7.574 3.694l-.851.639A2.492 2.492 0 0 1 8 6.5C8 7.878 6.878 9 5.5 9S3 7.878 3 6.5c0-.932.519-1.737 1.277-2.167l-.851-.639A3.485 3.485 0 0 0 2 6.5a3.5 3.5 0 1 0 7 0a3.485 3.485 0 0 0-1.426-2.806zM7.5 2.5L6.5 1h-2l-1 1.5l2 1.5l2-1.5z" fill="currentColor"/>`),
		g.Group(children),
	)
}