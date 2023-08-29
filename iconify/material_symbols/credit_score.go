package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreditScore(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.125 20H4q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v6h-2.725l-4.325 4.325l-2.85-2.8l-4.225 4.225l2.25 2.25Zm4.825 2l-4.25-4.25l1.4-1.4l2.85 2.8l5.65-5.65l1.4 1.45L14.95 22ZM4 8v4h16V8H4Z"/>`),
		g.Group(children),
	)
}