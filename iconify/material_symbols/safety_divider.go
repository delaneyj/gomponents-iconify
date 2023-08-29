package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SafetyDivider(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19V5h2v14h-2ZM1 16v-.575q0-.6.325-1.1t.9-.75q.65-.275 1.338-.425T5 13q.75 0 1.438.15t1.337.425q.575.25.9.75t.325 1.1V16H1Zm14 0v-.575q0-.6.325-1.1t.9-.75q.65-.275 1.337-.425T19 13q.75 0 1.437.15t1.338.425q.575.25.9.75t.325 1.1V16h-8ZM5 12q-.825 0-1.413-.588T3 10q0-.825.588-1.413T5 8q.825 0 1.413.588T7 10q0 .825-.588 1.413T5 12Zm14 0q-.825 0-1.413-.588T17 10q0-.825.588-1.413T19 8q.825 0 1.413.588T21 10q0 .825-.588 1.413T19 12Z"/>`),
		g.Group(children),
	)
}