package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WoundsInjuries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.4 22L2 20.6l2.2-2.2l-2.2-.6l.525-1.925l2.175.6l-.575-2.2l1.925-.525l.6 2.2l3.425-3.425l-2.2-.6L8.4 10l2.175.575L10 8.4l1.925-.525l.6 2.2L15.95 6.65l-2.2-.6l.525-1.925l2.2.575l-.6-2.175L17.8 2l.6 2.2L20.6 2L22 3.4l-2.2 2.2l2.2.6l-.525 1.925l-2.175-.6l.575 2.2l-1.925.525l-.6-2.2l-3.425 3.425l2.2.6L15.6 14l-2.175-.575L14 15.6l-1.925.525l-.6-2.2L8.05 17.35l2.2.6l-.525 1.925l-2.2-.575l.6 2.175L6.2 22l-.6-2.2L3.4 22Z"/>`),
		g.Group(children),
	)
}