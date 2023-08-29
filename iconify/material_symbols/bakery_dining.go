package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BakeryDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.875 17.7l-2.475-.85l1.95-5.375l2.125 4.375q.35.675-.3 1.375t-1.3.475Zm-6.275-.95l1.075-9.35q.05-.375.337-.513t.688-.012l2.5.925q.35.15.463.45t-.013.65l-2.825 7.85H14.6Zm-7.4 0L4.375 8.9q-.125-.35-.013-.663t.463-.437l2.5-.925q.35-.15.663-.012t.362.537l1.05 9.35H7.2Zm-4.3.95q-.65.2-1.175-.488T1.55 15.85l2.15-4.375l1.925 5.375l-2.725.85Zm8-.95L9.7 6q-.05-.425.238-.713T10.65 5h2.7q.425 0 .713.288T14.3 6l-1.2 10.75h-2.2Z"/>`),
		g.Group(children),
	)
}