package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldPerson(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13q1.45 0 2.475-1.025T15.5 9.5q0-1.45-1.025-2.475T12 6q-1.45 0-2.475 1.025T8.5 9.5q0 1.45 1.025 2.475T12 13Zm0 9q-3.65-.925-5.825-4T4 11.1V5l8-3l8 3v6.1q0 3.825-2.175 6.9T12 22Zm0-2.1q1.475-.475 2.613-1.488t1.987-2.287q-1.075-.55-2.238-.838T12 15q-1.2 0-2.363.288t-2.237.837q.85 1.275 1.988 2.288T12 19.9Z"/>`),
		g.Group(children),
	)
}