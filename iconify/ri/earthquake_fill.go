package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarthquakeFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.327 1.612a1 1 0 0 1 1.246-.08l.1.08L23 11h-3v9a1 1 0 0 1-.883.993L19 21h-6.5l2.5-4l-3.5-3l4-3L13 9l.5-3l-3 3l2.5 2l-5 3l3.75 3.5L8.5 21H5a1 1 0 0 1-.993-.883L4 20v-9H1l10.327-9.388Z"/>`),
		g.Group(children),
	)
}