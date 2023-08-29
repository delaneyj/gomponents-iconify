package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanDeleteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9h5l-5-5v5Zm2.9 12.5l-1.4-1.4l2.1-2.1l-2.1-2.1l1.4-1.4l2.1 2.1l2.1-2.1l1.4 1.4l-2.075 2.1l2.075 2.1l-1.4 1.4l-2.1-2.075l-2.1 2.075ZM4 22V2h10l6 6v4.35q-.475-.175-.975-.262T18 12q-2.5 0-4.25 1.738T12 17.975q0 1.125.4 2.163T13.55 22H4Z"/>`),
		g.Group(children),
	)
}