package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningCableOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M4.5 5.5h6m-6 0a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1m-6 0v-4a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4M9 2.5H6M5.5 13v2m4-2v2"/>`),
		g.Group(children),
	)
}