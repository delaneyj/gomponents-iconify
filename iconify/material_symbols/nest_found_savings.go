package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestFoundSavings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l-3-3H5q-.825 0-1.413-.588T3 18V4q0-.825.588-1.413T5 2h14q.825 0 1.413.588T21 4v14q0 .825-.588 1.413T19 20h-4l-3 3Zm0-7q2.15 0 3.575-1.5T17 11V6h-5Q9.975 6 8.487 7.425T7 11q0 .75.213 1.425t.587 1.25l-.4.4q-.3.3-.3.7t.3.7q.3.3.688.288t.712-.288l.375-.35q.6.425 1.313.65T12 16Zm1.25-6.375q.3.3.3.713t-.3.712l-1.875 1.875q-.3.3-.713.3t-.712-.3q-.275-.3-.275-.713t.275-.687l1.9-1.9q.3-.3.7-.3t.7.3Z"/>`),
		g.Group(children),
	)
}