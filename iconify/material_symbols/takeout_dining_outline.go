package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TakeoutDiningOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.8 18h8.45l.475-7H7.3l.5 7Zm-.65-9h9.725l.075-1.25L14.15 5h-4.3l-2.8 2.75l.1 1.25Zm-1.9 1.7L2 7.45l1.4-1.4L5 7.65l-.05-.6L9 3h6l4.05 4.05l-.05.6l1.6-1.6l1.4 1.4l-3.25 3.25H5.25Zm.7 9.3l-.7-9.3h13.5l-.7 9.3H5.95ZM12 9Zm.025 2Z"/>`),
		g.Group(children),
	)
}