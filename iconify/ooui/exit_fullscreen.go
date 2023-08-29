package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExitFullscreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 7V1H5v4H1v2zM5 19h2v-6H1v2h4zm10-4h4v-2h-6v6h2zm0-8h4V5h-4V1h-2v6z"/>`),
		g.Group(children),
	)
}