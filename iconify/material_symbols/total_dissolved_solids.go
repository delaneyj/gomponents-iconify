package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TotalDissolvedSolids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.3 11.8q.625-2.175 2.538-4.612T12 2q3.25 2.75 5.163 5.188T19.7 11.8H4.3Zm15.55 3.65q-.225 1.225-.788 2.313t-1.437 1.962q-.875.875-1.963 1.425t-2.287.75l6.475-6.45Zm-4.075-1.65h2.85l-8.1 8.1q-.575-.125-1.113-.287t-1.037-.438l7.4-7.375Zm-5.725 0h2.85l-6.225 6.225q-.375-.35-.712-.725t-.613-.8l4.7-4.7Zm-5.75 0h2.85L4.4 16.55q-.2-.675-.238-1.375T4.3 13.8Z"/>`),
		g.Group(children),
	)
}