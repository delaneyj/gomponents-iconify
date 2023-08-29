package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.708 15.44a6.968 6.968 0 0 0 3.997 1.266a7 7 0 1 0 6.59-9.413A7 7 0 1 0 4.708 15.44Zm1.147-1.64a4.976 4.976 0 0 0 2.432.886a6.97 6.97 0 0 1 1.256-4.408a6.97 6.97 0 0 1 3.713-2.687a5 5 0 1 0-7.4 6.21Zm12.29-3.603a4.977 4.977 0 0 0-2.432-.885a6.97 6.97 0 0 1-1.256 4.408a6.97 6.97 0 0 1-3.713 2.687a5 5 0 1 0 7.4-6.21Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}