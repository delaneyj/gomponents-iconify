package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveFileRenameOutlineOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 21l4-4h8v4Zm-6-2h1.4l8.625-8.625l-1.4-1.4L4 17.6ZM18.3 8.925l-4.25-4.2L16.875 1.9L21.1 6.125ZM2 21v-4.25l10.6-10.6l4.25 4.25L6.25 21ZM13.325 9.675l-.7-.7l1.4 1.4Z"/>`),
		g.Group(children),
	)
}