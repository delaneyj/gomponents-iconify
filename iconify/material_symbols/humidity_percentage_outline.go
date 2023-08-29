package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HumidityPercentageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.5 18q.625 0 1.063-.438T16 16.5q0-.625-.438-1.063T14.5 15q-.625 0-1.063.438T13 16.5q0 .625.438 1.063T14.5 18Zm-5.05-.05l6.5-6.5l-1.4-1.4l-6.5 6.5l1.4 1.4ZM9.5 13q.625 0 1.063-.438T11 11.5q0-.625-.438-1.063T9.5 10q-.625 0-1.063.438T8 11.5q0 .625.438 1.063T9.5 13Zm2.5 9q-3.425 0-5.713-2.35T4 13.8q0-2.5 1.988-5.438T12 2q4.025 3.425 6.013 6.363T20 13.8q0 3.5-2.288 5.85T12 22Zm0-2q2.6 0 4.3-1.763T18 13.8q0-1.825-1.513-4.125T12 4.65Q9.025 7.375 7.513 9.675T6 13.8q0 2.675 1.7 4.438T12 20Zm0-8Z"/>`),
		g.Group(children),
	)
}