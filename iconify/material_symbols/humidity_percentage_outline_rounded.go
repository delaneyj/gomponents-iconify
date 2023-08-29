package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumidityPercentageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 18q.625 0 1.063-.438T16 16.5q0-.625-.438-1.063T14.5 15q-.625 0-1.063.438T13 16.5q0 .625.438 1.063T14.5 18Zm-5.75-.75q.3.3.7.3t.7-.3l5.1-5.1q.3-.3.3-.7t-.3-.7q-.3-.3-.713-.3t-.712.3L8.75 15.825q-.3.3-.3.713t.3.712ZM9.5 13q.625 0 1.063-.438T11 11.5q0-.625-.438-1.063T9.5 10q-.625 0-1.063.438T8 11.5q0 .625.438 1.063T9.5 13Zm2.5 9q-3.425 0-5.713-2.35T4 13.8q0-1.55.7-3.1t1.75-2.975Q7.5 6.3 8.725 5.05T11 2.875q.2-.2.463-.287T12 2.5q.275 0 .537.088t.463.287q1.05.925 2.275 2.175t2.275 2.675Q18.6 9.15 19.3 10.7t.7 3.1q0 3.5-2.288 5.85T12 22Zm0-2q2.6 0 4.3-1.763T18 13.8q0-1.825-1.513-4.125T12 4.65Q9.025 7.375 7.513 9.675T6 13.8q0 2.675 1.7 4.438T12 20Zm0-8Z"/>`),
		g.Group(children),
	)
}