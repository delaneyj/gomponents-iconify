package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18c0 .7.12 1.36.34 2H12c-3.31 0-6-2.69-6-6c0-4 6-10.75 6-10.75S16.31 8.1 17.62 12c-.69.06-1.34.22-1.95.47C15 10.68 13.5 8.33 12 6.39C10 8.96 8 12.23 8 14c0 2.21 1.79 4 4 4m7-1v-3h-2v3h-3v2h3v3h2v-3h3v-2h-3Z"/>`),
		g.Group(children),
	)
}