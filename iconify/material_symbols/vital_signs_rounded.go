package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VitalSignsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 20q-.475 0-.85-.275t-.55-.7L5.3 13H2q-.425 0-.713-.288T1 12q0-.425.288-.713T2 11h4q.325 0 .563.175t.362.475L9 17.1l4.6-12.125q.175-.425.55-.7T15 4q.475 0 .85.275t.55.7L18.7 11H22q.425 0 .713.288T23 12q0 .425-.288.713T22 13h-4q-.325 0-.563-.175t-.362-.475L15 6.9l-4.6 12.125q-.175.425-.55.7T9 20Z"/>`),
		g.Group(children),
	)
}