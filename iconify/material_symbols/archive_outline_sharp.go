package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 18l4-4l-1.4-1.4l-1.6 1.6V10h-2v4.2l-1.6-1.6L8 14l4 4ZM5 8v11h14V8H5ZM3 21V5.8L5.3 3h13.4L21 5.8V21H3ZM5.4 6h13.2l-.85-1H6.25L5.4 6Zm6.6 7.5Z"/>`),
		g.Group(children),
	)
}