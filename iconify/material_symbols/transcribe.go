package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Transcribe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.05 15.95q-1.4-1.325-2.225-3.125T15 9q0-2.025.825-3.825T18.05 2.05l1.55 1.6q-1.1 1.025-1.725 2.4T17.25 9q0 1.575.625 2.975T19.6 14.4l-1.55 1.55Zm3.2-3.2q-.8-.725-1.25-1.688T19.55 9q0-1.1.45-2.063t1.25-1.687l1.6 1.6q-.45.425-.725.963T21.85 9q0 .65.275 1.188t.725.962l-1.6 1.6ZM9 13q-1.65 0-2.825-1.175T5 9q0-1.65 1.175-2.825T9 5q1.65 0 2.825 1.175T13 9q0 1.65-1.175 2.825T9 13Zm-8 8v-2.8q0-.825.425-1.55t1.175-1.1q1.275-.65 2.875-1.1T9 14q1.925 0 3.525.45t2.875 1.1q.75.375 1.175 1.1T17 18.2V21H1Z"/>`),
		g.Group(children),
	)
}