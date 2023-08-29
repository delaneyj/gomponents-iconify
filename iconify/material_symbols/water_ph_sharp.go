package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterPhSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 22v-6h5v4h-3.5v2H11Zm6 0v-6h1.5v2h2v-2H22v6h-1.5v-2.5h-2V22H17Zm-4.5-3.5h2v-1h-2v1ZM9 21.95q-3.05-.35-5.025-2.625T2 13.8q0-2.5 1.988-5.438T10 2q4.025 3.425 6.013 6.363T18 13.8v.2H9v7.95Z"/>`),
		g.Group(children),
	)
}