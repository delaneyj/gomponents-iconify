package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HotelClassOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.7 15.5l3.8-3.25l3 .25l-4.4 3.825L20.4 22l-2.55-1.55l-1.15-4.95Zm-2.35-7.3L13.3 5.75L14.45 3l2.3 5.425l-2.4-.225Zm-7 9.625l3.15-1.9l3.15 1.925l-.825-3.6l2.775-2.4l-3.65-.325l-1.45-3.4L9.05 11.5l-3.65.325l2.775 2.425l-.825 3.575ZM4.325 22l1.625-7.025L.5 10.25l7.2-.625L10.5 3l2.8 6.625l7.2.625l-5.45 4.725L16.675 22L10.5 18.275L4.325 22Zm6.175-8.5Z"/>`),
		g.Group(children),
	)
}