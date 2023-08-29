package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BakeryDiningOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.025 18Zm8.075-1.05l.85-.85l-1.45-2.7l-1.05 2.7l1.65.85Zm-5-.95h1.2l2.55-6.325l-3.1-1.225L15.1 16Zm-7.4 0h1.2l-.65-7.55l-3.1 1.225L7.7 16Zm-3.8.95l1.65-.85l-1.05-2.7l-1.45 2.7l.85.85Zm7-.95h2.2l.8-9h-3.8l.8 9Zm-7.375 3.425l-2.9-2.9L3.5 11l-.95-2.425l5.5-2.2L7.925 5h8.125l-.1 1.375l5.5 2.2l-.95 2.375l2.9 5.55l-2.875 2.875L17.7 18H6.25l-2.725 1.425Z"/>`),
		g.Group(children),
	)
}