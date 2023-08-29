package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAssistantDirection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.5 10H8v5h2v-3h3.5v2.5L17 11l-3.5-3.5V10zM12 1C5.9 1 1 5.9 1 12s4.9 11 11 11s11-4.9 11-11S18.1 1 12 1zm8.31 11l-8.34 8.37L3.62 12l8.34-8.37L20.31 12z"/>`),
		g.Group(children),
	)
}