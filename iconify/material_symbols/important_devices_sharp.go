package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImportantDevicesSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21v-9.95h6V21h-6Zm1-2h4v-5.95h-4V19Zm3-16v6.05h-8.1l-.9-2.8l-.9 2.8H7.25l2.3 1.85l-.85 2.85L11 12l2.3 1.75l-.85-2.85L14 9.65V17h-2v2h2v2H8v-2h2v-2H2V3h18Z"/>`),
		g.Group(children),
	)
}