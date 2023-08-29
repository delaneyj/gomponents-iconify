package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DriveFileRenameOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 21l4-4h8v4H10Zm8.3-12.075l-4.25-4.2l1.4-1.4q.575-.575 1.413-.575t1.412.575l1.4 1.4q.575.575.6 1.388t-.55 1.387L18.3 8.925ZM16.85 10.4L6.25 21H2v-4.25l10.6-10.6l4.25 4.25Z"/>`),
		g.Group(children),
	)
}