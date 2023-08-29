package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Organization(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#90CAF9" d="M42 42H6V10c0-2.2 1.8-4 4-4h28c2.2 0 4 1.8 4 4v32z"/><path fill="#64B5F6" d="M6 42h36v2H6z"/><path fill="#1565C0" d="M31 27h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20 8h6v5h-6zm-20 0h6v5h-6zm20-16h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm20-8h6v5h-6zm-10 0h6v5h-6zm-10 0h6v5h-6zm10 24h6v9h-6z"/>`),
		g.Group(children),
	)
}