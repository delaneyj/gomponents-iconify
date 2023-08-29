package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15q-1.275 0-2.138-.863T16 12q0-1.275.863-2.138T19 9q1.275 0 2.138.863T22 12q0 1.275-.863 2.138T19 15Zm-8.825-2H3q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h7.175L8.3 9.1q-.275-.275-.288-.688T8.3 7.7q.275-.275.7-.275t.7.275l3.6 3.6q.3.3.3.7t-.3.7l-3.6 3.6q-.3.3-.7.288t-.7-.313q-.275-.3-.288-.7t.288-.7L10.175 13Z"/>`),
		g.Group(children),
	)
}