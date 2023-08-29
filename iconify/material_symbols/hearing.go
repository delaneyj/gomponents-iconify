package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hearing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22q-1.55 0-2.538-.775T12.95 18.95q-.425-1.25-.812-1.75t-1.788-1.6Q8.8 14.35 7.9 12.775T7 9q0-2.975 2.013-4.987T14 2q2.975 0 4.988 2.013T21 9h-2q0-2.125-1.438-3.563T14 4q-2.125 0-3.563 1.438T9 9q0 1.7.675 2.9t1.925 2.15q1.3.95 2.025 1.85t1.075 1.95q.35 1.1.838 1.625T17 20q.825 0 1.413-.588T19 18h2q0 1.65-1.175 2.825T17 22ZM6.2 16.75q-1.475-1.5-2.338-3.488T3 9q0-2.3.863-4.3T6.2 1.2l1.45 1.4Q6.4 3.85 5.7 5.488T5 9q0 1.85.7 3.475t1.95 2.875l-1.45 1.4ZM14 11.5q-1.05 0-1.775-.738T11.5 9q0-1.05.725-1.775T14 6.5q1.05 0 1.775.725T16.5 9q0 1.025-.725 1.763T14 11.5Z"/>`),
		g.Group(children),
	)
}