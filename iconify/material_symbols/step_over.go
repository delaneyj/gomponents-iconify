package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepOver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19q-1.25 0-2.125-.875T9 16q0-1.25.875-2.125T12 13q1.25 0 2.125.875T15 16q0 1.25-.875 2.125T12 19Zm-7.925-7q.35-2.975 2.6-4.988T11.975 5q1.825 0 3.375.738T18 7.75V5h2v7h-7v-2h4.2q-.8-1.35-2.163-2.175T12 7Q9.8 7 8.125 8.425T6.1 12H4.075Z"/>`),
		g.Group(children),
	)
}