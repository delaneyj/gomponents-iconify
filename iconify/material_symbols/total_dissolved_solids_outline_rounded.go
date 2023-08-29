package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TotalDissolvedSolidsOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.3 11.8q.775-2.65 2.913-5.063t4.137-4.162q.275-.25.65-.25t.65.25q2 1.75 4.138 4.163T19.7 11.8h-2.1q-.575-1.55-1.975-3.338T12 4.65Q9.8 6.675 8.387 8.463T6.4 11.8H4.3Zm15.55 3.65q-.225 1.225-.788 2.313t-1.437 1.962q-.875.875-1.963 1.425t-2.287.75l6.475-6.45Zm-4.075-1.65h2.85l-8.1 8.1q-.575-.125-1.113-.287t-1.037-.438l7.4-7.375Zm-5.725 0h2.85l-6.225 6.225q-.375-.35-.712-.725t-.613-.8l4.7-4.7Zm-6.075 0H7.15L4.4 16.55q-.2-.675-.3-1.363T3.975 13.8Zm8.025-2Z"/>`),
		g.Group(children),
	)
}