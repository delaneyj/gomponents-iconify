package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func O(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M285 113c152 0 275 121 275 271S437 654 285 654S0 534 0 384s133-271 285-271zm-1 474c114 0 207-90 207-203s-93-204-207-204c-115 0-215 91-215 204s100 203 215 203z"/>`),
		g.Group(children),
	)
}