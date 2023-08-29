package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOneFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9.5 5.5v11L13 22.421V11.423l-9.502-5.5L12 1ZM2.5 7.655V17.5l8.5 4.921v-9.845l-8.5-4.92Z"/>`),
		g.Group(children),
	)
}