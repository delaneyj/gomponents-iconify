package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileOpenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22q-.825 0-1.413-.588T4 20V4q0-.825.588-1.413T6 2h8l6 6v6h-2V9h-5V4H6v16h9v2H6Zm15.95.375L19 19.425v2.225h-2V16h5.65v2H20.4l2.95 2.95l-1.4 1.425ZM6 20V4v16Z"/>`),
		g.Group(children),
	)
}