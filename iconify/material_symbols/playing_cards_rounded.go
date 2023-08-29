package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayingCardsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.375 14.125l.775-2.8q.1-.3-.013-.6t-.387-.5L13.375 8.6q-.2-.15-.438-.062t-.312.337l-.775 2.8q-.1.3.013.6t.387.5l2.375 1.625q.2.15.437.063t.313-.338ZM4 18.825l-.825-.4q-.775-.325-1.05-1.113T2.2 15.75l1.8-3.9v6.975ZM8 21q-.825 0-1.413-.6T6 18.975V13l2.675 7.35q.075.175.125.338t.175.312H8Zm5.15-.125q-.775.275-1.55-.075t-1.05-1.125L6.125 7.45q-.275-.775.075-1.538t1.125-1.037l7.525-2.75q.775-.275 1.538.075t1.037 1.125l4.45 12.225q.275.775-.075 1.538t-1.125 1.037l-7.525 2.75Z"/>`),
		g.Group(children),
	)
}