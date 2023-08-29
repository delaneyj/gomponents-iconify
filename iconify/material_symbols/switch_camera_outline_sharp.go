package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchCameraOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V5h5.15L9 3h6l1.85 2H22v16H2Zm2-2h16V7h-4.05l-1.825-2h-4.25L8.05 7H4v12Zm8-6Zm-3 4l1.4-1.4l-1.55-1.55h6.3L13.6 15.6L15 17l4-4l-4-4l-1.4 1.4l1.6 1.65H8.8l1.6-1.65L9 9l-4 4l4 4Z"/>`),
		g.Group(children),
	)
}