package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoodBankOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Zm0-2h12v-9l-6-4.5L6 10v9Zm6-6.75ZM10.5 18q.2 0 .35-.15t.15-.35V14q.625 0 1.063-.438T12.5 12.5V10q0-.2-.15-.35T12 9.5q-.2 0-.35.15t-.15.35v2.5H11V10q0-.2-.15-.35t-.35-.15q-.2 0-.35.15T10 10v2.5h-.5V10q0-.2-.15-.35T9 9.5q-.2 0-.35.15T8.5 10v2.5q0 .625.438 1.063T10 14v3.5q0 .2.15.35t.35.15Zm4 0q.2 0 .35-.15t.15-.35v-7.35q0-.275-.188-.463T14.35 9.5q-.675 0-1.012.625T13 11.5v2q0 .4.363.563T14 14.5v3q0 .2.15.35t.35.15Z"/>`),
		g.Group(children),
	)
}