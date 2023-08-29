package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnergySavingsLeaf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-1.6 0-3.025-.525T6.4 19.025L4.425 21L3 19.575L4.975 17.6q-.925-1.15-1.45-2.588T3 12q0-3.75 2.638-6.375T12 3h9v9q0 3.725-2.625 6.363T12 21Zm-1.325-4.125l5.175-4.625q.25-.225.15-.537t-.45-.363l-4.05-.4l2.425-3.325q.1-.125.088-.262T13.9 7.1q-.125-.125-.288-.125t-.287.125l-5.15 4.625q-.25.225-.15.538t.45.362l4.05.4l-2.45 3.325q-.075.125-.075.263t.125.262q.125.125.275.125t.275-.125Z"/>`),
		g.Group(children),
	)
}