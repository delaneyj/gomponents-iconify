package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HideImageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 18.15L5.85 3H19q.825 0 1.413.588T21 5v13.15ZM5 21q-.825 0-1.413-.588T3 19V5.8l-.9-.9q-.275-.275-.275-.7t.275-.7q.275-.275.7-.275t.7.275l17 17q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-.9-.9H5Zm9.175-4l-2.1-2.1l-.825 1.1l-1.85-2.475q-.15-.2-.4-.2t-.4.2l-2 2.675q-.2.25-.05.525T7 17h7.175Z"/>`),
		g.Group(children),
	)
}