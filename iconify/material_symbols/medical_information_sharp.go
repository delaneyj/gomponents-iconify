package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicalInformationSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 18h2v-2h2v-2H9v-2H7v2H5v2h2v2Zm6-3.5h6V13h-6v1.5Zm0 3h4V16h-4v1.5ZM2 22V7h7V2h6v5h7v15H2Zm9-13h2V4h-2v5Z"/>`),
		g.Group(children),
	)
}