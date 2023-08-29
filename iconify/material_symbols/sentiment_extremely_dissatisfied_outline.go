package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SentimentExtremelyDissatisfiedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-2q3.325 0 5.663-2.337T20 12q0-3.325-2.337-5.663T12 4Q8.675 4 6.337 6.337T4 12q0 3.325 2.337 5.663T12 20Zm1.5-9.1l.5-.3q.05.6.475 1t1.025.4q.625 0 1.063-.438T17 10.5q0-.375-.175-.713t-.475-.537l.65-.375L16.5 8L13 10l.5.9Zm-3 0l.5-.9l-3.5-2l-.5.875l.65.375q-.3.2-.475.537T7 10.5q0 .625.438 1.063T8.5 12q.6 0 1.025-.4t.475-1l.5.3ZM12 13q-1.775 0-3.125 1.137T6.975 17h10.05q-.55-1.725-1.9-2.863T12 13Zm0-1Z"/>`),
		g.Group(children),
	)
}