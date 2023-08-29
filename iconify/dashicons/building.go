package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Building(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 20h14V0H3v20zM7 3H5V1h2v2zm4 0H9V1h2v2zm4 0h-2V1h2v2zM7 6H5V4h2v2zm4 0H9V4h2v2zm4 0h-2V4h2v2zM7 9H5V7h2v2zm4 0H9V7h2v2zm4 0h-2V7h2v2zm-8 3H5v-2h2v2zm4 0H9v-2h2v2zm4 0h-2v-2h2v2zm-4 7H5v-6h6v6zm4-4h-2v-2h2v2zm0 3h-2v-2h2v2z"/>`),
		g.Group(children),
	)
}