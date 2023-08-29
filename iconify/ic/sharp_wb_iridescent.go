package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpWbIridescent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 15h14V9.05H5V15zm6-14v3h2V1h-2zm8.04 2.6l-1.79 1.79l1.41 1.41l1.8-1.79l-1.42-1.41zM13 23v-2.95h-2V23h2zm7.45-3.91l-1.8-1.79l-1.41 1.41l1.79 1.8l1.42-1.42zM3.55 5.01L5.34 6.8l1.41-1.41L4.96 3.6L3.55 5.01zM4.96 20.5l1.79-1.8l-1.41-1.41l-1.79 1.79l1.41 1.42z"/>`),
		g.Group(children),
	)
}