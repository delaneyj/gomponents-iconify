package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InstallMobileOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h7v2H7v1h7v2H7v12h10v-2h2v5q0 .825-.588 1.413T17 23H7Zm0-3v1h10v-1H7Zm10-9.8V4q0-.425.288-.713T18 3q.425 0 .713.288T19 4v6.2l1.9-1.9q.275-.275.7-.275t.7.275q.275.275.275.7t-.275.7l-3.6 3.6q-.3.3-.7.3t-.7-.3l-3.6-3.6q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l1.9 1.9ZM7 4V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}