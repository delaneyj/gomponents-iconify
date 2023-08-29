package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsMotorsportsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12Zm2 6q2.5 0 4.25-1.75T20 12q0-2.525-1.838-4.263T13.75 6q-1.2 0-2.325.275T9.25 7.1l2.5 1q1.025.425 1.638 1.313T14 11.4q0 1.5-1.038 2.55T10.45 15h-6.4Q4 15.6 4 16.363V18h10Zm-9.6-5h6q.675 0 1.137-.463T12 11.4q0-.475-.263-.863T11 9.95l-3.7-1.5q-1.05.925-1.788 2.1T4.4 13Zm9.6 7H4q-.825 0-1.413-.588T2 18v-2.25q0-2.45.925-4.588t2.513-3.725Q7.024 5.85 9.175 4.925T13.75 4q1.7 0 3.2.625t2.625 1.713Q20.7 7.425 21.35 8.875T22 12q0 1.65-.625 3.113t-1.713 2.55q-1.087 1.087-2.55 1.712T14 20Z"/>`),
		g.Group(children),
	)
}