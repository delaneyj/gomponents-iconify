package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TakeoutDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.625 10L2 7.45l1.4-1.4L5 7.65l-.05-.6L9 3h6l4.05 4.05l-.05.6l1.6-1.6l1.4 1.4L19.375 10H4.625ZM5.95 20l-.65-8.45h13.4L18.05 20H5.95Z"/>`),
		g.Group(children),
	)
}