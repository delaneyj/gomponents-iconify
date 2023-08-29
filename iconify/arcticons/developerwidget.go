package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Developerwidget(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 34.972V40.9a2.608 2.608 0 0 0 2.6 2.6h20.8a2.608 2.608 0 0 0 2.6-2.6V7.1a2.608 2.608 0 0 0-2.6-2.6H13.6A2.608 2.608 0 0 0 11 7.1v5.928"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.926 31.947L20.873 24l-7.947-7.947m7.947 15.894h8.857"/>`),
		g.Group(children),
	)
}