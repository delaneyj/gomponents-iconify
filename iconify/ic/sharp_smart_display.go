package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSmartDisplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 4H2v16h20V4zM9.5 16.5v-9l7 4.5l-7 4.5z"/>`),
		g.Group(children),
	)
}