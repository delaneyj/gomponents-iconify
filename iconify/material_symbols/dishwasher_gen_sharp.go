package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DishwasherGenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 7q.425 0 .713-.288T9 6q0-.425-.288-.713T8 5q-.425 0-.713.288T7 6q0 .425.288.713T8 7Zm3 0q.425 0 .713-.288T12 6q0-.425-.288-.713T11 5q-.425 0-.713.288T10 6q0 .425.288.713T11 7Zm1 11q1.575 0 2.688-1.075T15.8 14.3q0-.725-.25-1.413T14.8 11.7L12 8.9l-2.7 2.7q-.55.55-.838 1.25T8.2 14.3q.05 1.55 1.15 2.625T12 18Zm0-1.9q-.75 0-1.275-.525T10.2 14.3q0-.375.138-.713t.412-.612l1.25-1.25l1.225 1.225q.275.275.425.625t.15.725q0 .75-.525 1.275T12 16.1ZM4 22V2h16v20H4Z"/>`),
		g.Group(children),
	)
}