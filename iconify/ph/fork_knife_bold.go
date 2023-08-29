package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkKnifeBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M68 88V40a12 12 0 0 1 24 0v48a12 12 0 0 1-24 0Zm152-48v184a12 12 0 0 1-24 0v-44h-44a12 12 0 0 1-12-12a273.23 273.23 0 0 1 7.33-57.82c10.09-41.76 29.43-69.85 55.94-81.18A12 12 0 0 1 220 40Zm-24 22.92C182.6 77 175 98 170.77 115.38a254.41 254.41 0 0 0-6.22 40.62H196ZM128 39a12 12 0 0 0-24 2l4 47.46a28 28 0 0 1-56 0L56 41a12 12 0 1 0-24-2l-4 48v1a52.1 52.1 0 0 0 40 50.59V224a12 12 0 0 0 24 0v-85.41A52.1 52.1 0 0 0 132 88v-1Z"/>`),
		g.Group(children),
	)
}