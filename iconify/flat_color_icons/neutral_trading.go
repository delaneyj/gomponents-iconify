package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NeutralTrading(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="#1565C0"><path d="M43.4 13L35 20V6z"/><path d="M4 11h34v4H4z"/></g><path fill="#2196F3" d="M40 23h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19h-4zm-6 0h4v19H4z"/>`),
		g.Group(children),
	)
}