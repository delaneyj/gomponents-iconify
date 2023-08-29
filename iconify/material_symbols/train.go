package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Train(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15.5V6q0-1.325.688-2.113t1.812-1.2q1.125-.412 2.563-.55T12 2q1.65 0 3.113.138t2.55.55q1.087.412 1.712 1.2T20 6v9.5q0 1.475-1.012 2.488T16.5 19l1.5 1.5v.5h-2l-2-2h-4l-2 2H6v-.5L7.5 19q-1.475 0-2.488-1.012T4 15.5ZM6 10h5V7H6v3Zm7 0h5V7h-5v3Zm-4.5 6q.65 0 1.075-.425T10 14.5q0-.65-.425-1.075T8.5 13q-.65 0-1.075.425T7 14.5q0 .65.425 1.075T8.5 16Zm7 0q.65 0 1.075-.425T17 14.5q0-.65-.425-1.075T15.5 13q-.65 0-1.075.425T14 14.5q0 .65.425 1.075T15.5 16Z"/>`),
		g.Group(children),
	)
}