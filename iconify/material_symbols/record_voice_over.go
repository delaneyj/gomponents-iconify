package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordVoiceOver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.95 15.95L18.4 14.4q1.1-1.025 1.725-2.425T20.75 9q0-1.575-.625-2.95T18.4 3.65l1.55-1.6q1.4 1.325 2.225 3.125T23 9q0 2.025-.825 3.825T19.95 15.95Zm-3.2-3.2l-1.6-1.6q.45-.425.725-.963T16.15 9q0-.65-.275-1.188t-.725-.962l1.6-1.6q.8.725 1.25 1.688T18.45 9q0 1.1-.45 2.063t-1.25 1.687ZM9 13q-1.65 0-2.825-1.175T5 9q0-1.65 1.175-2.825T9 5q1.65 0 2.825 1.175T13 9q0 1.65-1.175 2.825T9 13Zm-8 8v-2.8q0-.825.425-1.55t1.175-1.1q1.275-.65 2.875-1.1T9 14q1.925 0 3.525.45t2.875 1.1q.75.375 1.175 1.1T17 18.2V21H1Z"/>`),
		g.Group(children),
	)
}