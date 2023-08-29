package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepIntoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-1.25 0-2.125-.875T9 19q0-1.25.875-2.125T12 16q1.25 0 2.125.875T15 19q0 1.25-.875 2.125T12 22Zm-1-11.825V3q0-.425.288-.713T12 2q.425 0 .713.288T13 3v7.175L14.875 8.3q.275-.275.688-.275t.712.275q.3.3.3.713t-.3.712L12.7 13.3q-.3.3-.7.3t-.7-.3L7.7 9.7q-.275-.275-.288-.687T7.7 8.3q.275-.275.7-.275t.7.275l1.9 1.875Z"/>`),
		g.Group(children),
	)
}