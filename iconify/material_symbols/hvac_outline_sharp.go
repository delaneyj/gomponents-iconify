package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HvacOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q2.5 0 4.25-1.75T18 12q0-2.5-1.75-4.25T12 6Q9.5 6 7.75 7.75T6 12q0 2.5 1.75 4.25T12 18Zm0-2q-.725 0-1.4-.263T9.375 15h5.25q-.55.475-1.225.738T12 16Zm-3.45-2q-.2-.35-.325-.725T8.05 12.5h7.9q-.05.4-.175.775T15.45 14h-6.9Zm-.5-2.5q.05-.4.175-.775T8.55 10h6.9q.2.35.325.725t.175.775h-7.9ZM9.375 9q.55-.475 1.225-.738T12 8q.725 0 1.4.262T14.625 9h-5.25ZM3 21V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}