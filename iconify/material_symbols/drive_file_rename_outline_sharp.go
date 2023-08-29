package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveFileRenameOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 21l4-4h8v4Zm8.3-12.075l-4.25-4.2L16.875 1.9L21.1 6.125ZM2 21v-4.25l10.6-10.6l4.25 4.25L6.25 21Z"/>`),
		g.Group(children),
	)
}