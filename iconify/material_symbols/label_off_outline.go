package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.45 15.6L17 14.15L18.55 12L15 7H9.85l-2-2H15q.5 0 .925.213t.725.637L21 12l-2.55 3.6Zm1.35 7l-3.85-3.85q-.225.125-.45.188T15 19H5q-.825 0-1.413-.588T3 17V7q0-.275.063-.5t.187-.45L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM9.575 12.4Zm3.85-1.825ZM14.2 17L5 7.8V17h9.2Z"/>`),
		g.Group(children),
	)
}