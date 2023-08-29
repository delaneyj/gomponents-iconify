package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Recycling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.2 9.2l2.225-3.675l-1.475-2.45q-.3-.5-.862-.5t-.863.5L5.775 7.15L9.2 9.2Zm9.675 6.8l-2.225-3.7l3.475-2l1.6 2.675q.275.425.3.95t-.225.975q-.25.5-.738.8T20 16h-1.125ZM16 23l-4-4l4-4v2h4.75l-1.45 2.9q-.275.5-.75.8t-1.05.3H16v2Zm-9.675-2q-.5 0-.913-.263T4.8 20.05q-.2-.4-.187-.838t.237-.812L5.7 17H10v4H6.325ZM3.85 18.15L2.225 14.9q-.225-.45-.213-.963t.288-.962l.4-.675L1 11.275L6.475 9.9l1.375 5.5l-1.725-1.05l-2.275 3.8Zm13.5-8.55l-5.475-1.375L13.6 7.2L10.475 2H14q.525 0 .988.263t.737.712l1.3 2.175l1.7-1.05l-1.375 5.5Z"/>`),
		g.Group(children),
	)
}