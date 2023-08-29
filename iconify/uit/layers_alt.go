package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.5 2h-13a.5.5 0 0 0-.5.5V7H5.5a.5.5 0 0 0-.5.5V12H2.5a.5.5 0 0 0-.5.5v9a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V19h4.5a.5.5 0 0 0 .5-.5V16h4.5a.5.5 0 0 0 .5-.5v-13a.5.5 0 0 0-.5-.5zM11 18.5V21H3v-8h8v5.5zm5-3V18h-4v-5.5a.5.5 0 0 0-.5-.5H6V8h10v7.5zm5-.5h-4V7.5a.5.5 0 0 0-.5-.5H9V3h12v12z"/>`),
		g.Group(children),
	)
}