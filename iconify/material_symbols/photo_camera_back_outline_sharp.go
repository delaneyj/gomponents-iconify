package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraBackOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V5h5.15L9 3h6l1.85 2H22v16H2Zm2-2h16V7h-4.05l-1.825-2h-4.25L8.05 7H4v12Zm8-6Zm-6 4h12l-3.75-5l-3 4L9 13l-3 4Z"/>`),
		g.Group(children),
	)
}