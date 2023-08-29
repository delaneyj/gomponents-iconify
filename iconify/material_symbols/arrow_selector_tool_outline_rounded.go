package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSelectorToolOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 13.75L9.975 11h4.25L8 6.1v7.65Zm7.15 7.625q-.575.275-1.15.063t-.85-.788l-3-6.45l-2.325 3.25q-.425.6-1.125.375t-.7-.95V4.05q0-.625.563-.9t1.062.125l10.1 7.95q.575.425.338 1.1T17.1 13h-4.2l2.975 6.375q.275.575.063 1.15t-.788.85ZM9.975 11Z"/>`),
		g.Group(children),
	)
}