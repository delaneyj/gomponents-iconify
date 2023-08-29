package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRestaurantOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.35 9H19.7l-.85-3H5.225L4.35 9Zm7.675-1.5ZM6.975 13h10.1l-.25-2H7.25l-.275 2ZM5.15 20q-.45 0-.738-.338t-.237-.787L5.25 11H3.025q-.5 0-.788-.4t-.162-.875l1.425-5q.1-.325.35-.525t.6-.2H19.6q.35 0 .6.2t.35.525l1.425 5q.125.475-.162.875t-.788.4h-2.2l1.05 7.875q.05.45-.237.788T18.9 20q-.375 0-.662-.238t-.338-.612L17.35 15H6.7l-.55 4.15q-.05.375-.338.613T5.15 20Z"/>`),
		g.Group(children),
	)
}