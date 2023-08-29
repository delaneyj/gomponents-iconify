package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddressBookSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M6 4v5H5v2h3V6h16v20H8v-3H6v5h20V4zm10 6c-2.2 0-4 1.8-4 4c0 1.113.477 2.117 1.219 2.844A5.041 5.041 0 0 0 11 21h2a3 3 0 0 1 6 0h2a5.041 5.041 0 0 0-2.219-4.156C19.523 16.117 20 15.114 20 14c0-2.2-1.8-4-4-4zM6 12v2H5v2h3v-4zm10 0c1.117 0 2 .883 2 2s-.883 2-2 2s-2-.883-2-2s.883-2 2-2zM6 17v2H5v2h3v-4z"/>`),
		g.Group(children),
	)
}