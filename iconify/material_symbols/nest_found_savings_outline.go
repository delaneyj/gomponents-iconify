package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestFoundSavingsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 23l-3-3H5q-.825 0-1.413-.588T3 18V4q0-.825.588-1.413T5 2h14q.825 0 1.413.588T21 4v14q0 .825-.588 1.413T19 20h-4l-3 3Zm-7-5h4.8l2.2 2.2l2.2-2.2H19V4H5v14Zm7-7Zm0 5q2.15 0 3.575-1.5T17 11V6h-5Q9.975 6 8.487 7.425T7 11q0 .75.213 1.425t.587 1.25l-.4.4q-.3.3-.3.7t.3.7q.3.3.688.288t.712-.288l.375-.35q.6.425 1.313.65T12 16Zm0-2q-.375 0-.713-.1t-.637-.25l2.6-2.6q.3-.3.3-.713t-.3-.712q-.3-.3-.7-.3t-.7.3l-2.6 2.6q-.125-.275-.188-.587T9 11q0-1.3.9-2.15T12 8h3v3q0 1.25-.875 2.125T12 14Zm-.4-2.725Z"/>`),
		g.Group(children),
	)
}